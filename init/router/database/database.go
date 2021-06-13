package database

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

var db *gorm.DB

//Init database接続
func Init() {
	//configからデータベースのプロバイダとパスを取得しOpenする
	var err error
	url := os.Getenv("DATABASE_URL")
	connection, err := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err = gorm.Open("postgres", connection)
	if err != nil {
		panic("データベースへの接続失敗")
	}
}

//GetDB return db connection pool
func GetDB() *gorm.DB {
	return db
}
