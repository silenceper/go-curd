package model

import (
    "github.com/silenceper/go-blog/blog/config"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "log"
)

var (
    db *sql.DB
    err error
)

func init(){
    db, err = sql.Open("mysql", config.CFG_DSN)
    if err != nil{
        log.Fatal(err)
    }
}
