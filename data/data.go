package data

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB
func OpenDb()  {
	/* dbname:="root:admin@tcp/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	sql.Open("mysql",dbname)
	fmt.Println() */

	
}
