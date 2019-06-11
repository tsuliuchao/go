package main

import (
 db "github.com/liuchao/project/database"
)

func main() {
 defer db.SqlDB.Close()
 router := initRouter()
 router.Run(":8000")
}