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

func GetDetail(id int64)map[string]interface{}{
    var title string
    var content string
    var utime int64
    
    err := db.QueryRow("SELECT id,title,content,utime FROM article where id=?",id).Scan(&id,&title,&content,&utime)
    if err != nil{
        log.Fatal(err)
    }
    m:=make(map[string]interface{})
    m["id"]=id
    m["title"]=title
    m["content"]=content
    m["utime"]=utime
    return m
}

func InsertArticle(data map[string]interface{}) (int64,error){
    stmt,err:=db.Prepare("insert into article(`title`,`content`,`uid`,`utime`) values(?,?,?,?)")
    if err!=nil {
        return 0,err
    }
    res,err:=stmt.Exec(data["title"],data["content"],data["uid"],data["utime"])
    if err!=nil{
        return 0,err
    }
    id,err:=res.LastInsertId()
    if err!=nil{
        return 0,err
    }
    return id,nil
}

func UpdateArticle(data map[string]interface{},id int64) error{
    stmt,err:=db.Prepare("update article set title=?,content=?,utime=? where id=?")
    if err!=nil{
        return err
    }
    _,err=stmt.Exec(data["title"],data["content"],data["utime"],id)
    return err
}

func DeleteById(id int64) error{
    stmt,err:=db.Prepare("delete from article where id=?")
    if err!=nil{
        return err
    }
    _,err=stmt.Exec(id)
    return err
}
