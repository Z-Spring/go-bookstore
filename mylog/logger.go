package mylog

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

type Config struct {
	TimeFormat string
	UTC        bool
	SkipPaths  []string
}

type LogInfo struct {
	Level     string `json:"level,omitempty"`
	TimeStamp string `json:"ts,omitempty"`
	Caller    string `json:"caller,omitempty"`
	Message   string `json:"msg,omitempty"`
	Method    string `json:"method,omitempty"`
	Status    string `json:"status,omitempty"`
	Query     string `json:"query,omitempty"`
	Ip        string `json:"ip,omitempty"`
	UserAgent string `json:"user_agent,omitempty"`
	Latency   string `json:"latency,omitempty"`
}

// GinLogWithConfig rewrite a logger,add more info to logger output
func GinLogWithConfig(logger *zap.Logger, conf *Config) gin.HandlerFunc {
	skipPaths := make(map[string]bool, len(conf.SkipPaths))
	for _, path := range conf.SkipPaths {
		skipPaths[path] = true
	}

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		if _, ok := skipPaths[path]; !ok {
			end := time.Now()
			latency := end.Sub(start)
			if conf.UTC {
				end = end.UTC()
			}

			if len(c.Errors) > 0 {
				// Append error field if this is an erroneous request.
				for _, e := range c.Errors.Errors() {
					logger.Error(e)
				}
			} else {
				fields := []zapcore.Field{
					zap.String("method", c.Request.Method),
					zap.Int("status", c.Writer.Status()),
					zap.String("query", query),
					zap.String("ip", c.ClientIP()),
					zap.String("user-agent", c.Request.UserAgent()),
					zap.Duration("latency", latency),
				}
				if conf.TimeFormat != "" {
					fields = append(fields, zap.String("time", end.Format(conf.TimeFormat)))
				}
				logger.Info(path, fields...)
			}
		}
	}
}

// TimeEncoder time format
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006.01.02 15:04:05"))
}

var Log2 = make(chan *os.File, 10)

// NewZapProduction config zap logger like timeEncoder,level upper and so on...
func NewZapProduction() (*zap.Logger, error) {
	//var ZapFile io.Writer
	//var sugaredLogger *zap.SugaredLogger
	/*fileName := "logs/" + time.Now().Format("2006.01.02") + ".log"
	File2, err := CreateFile("logs", fileName)
	if err != nil {
		return nil, err
	}*/
	//todo 这里读到buf也不是一个好选择
	//log2 := make(chan zapcore.WriteSyncer, 10)
	writer := zapcore.AddSync(os.Stdout)
	//Log2 <- os.Stdout
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = TimeEncoder
	// Level 大写
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, writer, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger, nil
}

func CreateFile(path string, file string) (*os.File, error) {
	if err := os.MkdirAll(path, 0755); err != nil {
		log.Fatalf("%v", "CreateFileError")
		return nil, err
	}
	fileName, err := os.Create(file)
	if err != nil {
		return nil, err
	}
	return fileName, err
}

func GinRecoveryWithConfig(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
