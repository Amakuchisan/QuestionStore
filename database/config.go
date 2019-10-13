package database

import (
	"log"
	// MySQL Driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	dbDriver = "mysql"

	// TODO: read from environment values
	dataSource = "tts:tts@tcp(mysql-container:3306)/tts?parseTime=true"
)

// DB -- refered handler
var DB *sqlx.DB

// Init -- 自動でinitされると困るから静的に呼び出す
func init() {
	DB = GetDBSession()
}

// GetDBSession is return db connection
func GetDBSession() *sqlx.DB {
	var err error
	// dbConf := getDBConfig()
	DB, err = sqlx.Open(dbDriver, dataSource)
	
	if err != nil {
		log.Fatal("failed to connect db", err)
	}
	return DB
}

// func getDBConfig() string {
// 	dbConf := fmt.Sprintf(
// 		"%s:%s@tcp(%s:3306)/%s?parseTime=true",
// 		os.Getenv("MYSQL_USER"),
// 		os.Getenv("MYSQL_PASSWORD"),
// 		os.Getenv("MYSQL_HOST"),
// 		os.Getenv("MYSQL_DATABASE"),
// 	)
// 	return dbConf
// }