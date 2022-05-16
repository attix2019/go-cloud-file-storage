package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init(){
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "123456"
	dbName := "cloud_storage"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil{
		fmt.Println("fail to get sql connection " + err.Error())
	}
	db.SetMaxOpenConns(1000)
	PingDB(db)
}

func DbConn() *sql.DB {
	return db
}

func PingDB(db *sql.DB) {
	err = db.Ping()
	if err != nil{
		fmt.Println("ping failed " + err.Error())
	}
}

