package es

import (
	"bookstore/mylog"
	"context"
	"errors"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/estransport"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func ConfigWithEs() error {
	var (
		CreateEsConnectionError = errors.New("创建连接失败")
		//OpenFileFailure         = errors.New("打开文件失败")
		//CreateIndexFailure      = errors.New("创建索引失败")
	)

	//client, err := elasticsearch7.NewDefaultClient()
	cfg := elasticsearch7.Config{
		Logger: &estransport.JSONLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		},
	}
	client, err := elasticsearch7.NewClient(cfg)
	if err != nil {
		return CreateEsConnectionError
	}

	/*	file, err := os.Open("logs/2022.05.15.json")
		if err != nil {
			return OpenFileFailure
		}
		data, err := io.ReadAll(file)*/

	s := `{"level":"INFO","ts":"2022.05.17 22:37:38","caller":"mylog/logger.go:60","msg":"/books/12","method":"GET","status":200,"query":"","ip":"::1","user-agent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36","latency":0.0090668} `

	var logInfo []string
	logInfo = append(logInfo, string(s))

	indexer, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         "index",
		Client:        client,
		FlushBytes:    int(5e+6),
		FlushInterval: 3 * time.Second,
	})
	for v := range mylog.Log2 {
		log.Println(v)
		all, _ := io.ReadAll(v)
		indexer.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action: "index",
				Body:   strings.NewReader(string(all)),
			},
		)
	}

	/*	res, err := client.Index(
			"index2",
			strings.NewReader(v),
			//esutil.NewJSONReader(data),
			client.Index.WithRefresh("true"))
		log.Println(res.StatusCode)
		if err != nil {
			log.Println("index err:", err)
		}
		res.Body.Close()*/

	/*const indexName = "index"
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         indexName,
		Client:        client,
		FlushBytes:    int(5e+6),
		FlushInterval: 30 * time.Second,
	})
	if err != nil {
		log.Printf("error creating the indexer : %s", err)
		return CreateIndexFailure
	}
	var countSuccessful uint64
	for _, v := range logInfo {
		log.Println("data: ", v)
		err = bi.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action:     "create",
				DocumentID: strconv.Itoa(1),
				Body:       strings.NewReader(v),
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					atomic.AddUint64(&countSuccessful, 1)
				},

				// OnFailure is called for each failed operation
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						log.Printf("ERROR: %s", err)
					} else {
						log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
	}

	if err != nil {
		return err
	}
	if err := bi.Close(context.Background()); err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}
	if res, err := client.Indices.Delete([]string{indexName}, client.Indices.Delete.WithIgnoreUnavailable(true)); err != nil || res.IsError() {
		log.Fatalf("Cannot delete index: %s", err)
	}
	res, err := client.Indices.Create(indexName)
	if err != nil {
		log.Fatalf("Cannot create index: %s", err)
	}
	if res.IsError() {
		log.Fatalf("Cannot create index: %s", res)
	}*/
	return nil
}

func TransmitLogToEs() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := ConfigWithEs()
		if err != nil {
			log.Printf("es error: %v", err)
		}
	}
}
