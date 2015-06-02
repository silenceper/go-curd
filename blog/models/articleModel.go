package model

import (
    "log"
)

func GetList() []map[string]interface{}{
    rows, err := db.Query("SELECT id,title,content,utime FROM article")
    if err != nil{
        log.Fatal(err)
    }
    list:=[]map[string]interface{}{}
    
    for rows.Next() {
        m:=make(map[string]interface{})
        var id int64
        var title string
        var content string
        var utime int64
        err=rows.Scan(&id,&title,&content,&utime)
        m["id"]=id
        m["title"]=title
        m["content"]=content
        m["utime"]=utime
        list=append(list,m)
    }
    return list
}
