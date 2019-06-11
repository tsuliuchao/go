package database

import (
 "database/sql"
 _ "github.com/go-sql-driver/mysql"
 "log"
)

var SqlDB *sql.DB

func init() {
 var err error
 SqlDB, err = sql.Open("mysql", "pc_dev:123456@tcp(10.112.170.55:3306)/test?parseTime=true")
 if err != nil {
  log.Fatal(err.Error())
 }
// 设置数据库连接池
 SqlDB.SetMaxIdleConns(20)
 SqlDB.SetMaxOpenConns(20)
 err = SqlDB.Ping()
 if err != nil {
  log.Fatal(err.Error())
 }

}

//# bind by query
//$ curl -X GET "localhost:8085/testing?name=appleboy&address=xyz"
//# bind by json
//$ curl -X GET localhost:8085/testing --data '{"name":"JJ", "address":"xyz"}' -H "Content-Type:application/json"