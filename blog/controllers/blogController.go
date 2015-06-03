package controller

import (
    "net/http"
    "../models"
    "io"
    "time"
    "strconv"
    //"fmt"
)

func IndexController(w http.ResponseWriter,r *http.Request){
    //io.WriteString(w,"index Controller") 
    list:=model.GetList()
    bytes,err:=Render("index/index.html",map[string]interface{}{
        "list":list,
    })
    if err!=nil{
        panic(err)
    }
    w.Write(bytes)
}


func NewController(w http.ResponseWriter,r *http.Request){
    if r.Method=="GET"{
        bytes,err:=Render("index/new.html",map[string]interface{}{})
        if err!=nil{
            panic(err)
        }
        w.Write(bytes)
        return 
    }
    //处理请求
    r.ParseForm()
    title:=r.Form.Get("title")
    content:=r.Form.Get("content")
    if title=="" || content==""{
       io.WriteString(w,"必填字段不能为空")
    }
    data:=make(map[string]interface{},4)
    data["title"]=title
    data["content"]=content
    data["uid"]=1
    data["utime"]=time.Now().Unix()
    id,err:=model.InsertArticle(data)
    if err!=nil{
        io.WriteString(w,"插入失败")
        return 
    }
    //跳转到相应文章页面
    w.Header().Add("Location","/detail?id="+strconv.FormatInt(id,10))
    w.WriteHeader(302)
}

func UpdateController(w http.ResponseWriter,r *http.Request){
    r.ParseForm()
    if r.Method=="GET"{
        //根据id获取详情
        id:=r.Form.Get("id")
        id_num,err:=strconv.Atoi(id)
        if err!=nil{
            //w.WriteHeader(404)
            //io.WriteString(w,"参数错误")
            http.Error(w,"参数错误",404)
            return 
        }

        article:=model.GetDetail(int64(id_num))
        bytes,err:=Render("index/update.html",map[string]interface{}{
            "article":article,
        })
        if err!=nil{
            panic(err)
        }
        w.Write(bytes)
        return 
    }

    //处理update
    title:=r.Form.Get("title")
    content:=r.Form.Get("content")
    id:=r.Form.Get("id")
    if title=="" || content=="" ||id==""{
       io.WriteString(w,"必填字段不能为空")
    }
    id_num,err:=strconv.Atoi(id)
    if err!=nil{
        http.Error(w,"参数错误",404)
        return 
    }

    data:=make(map[string]interface{},4)
    data["title"]=title
    data["content"]=content
    data["uid"]=1
    data["utime"]=time.Now().Unix()
    err=model.UpdateArticle(data,int64(id_num))
    if err!=nil{
        http.Error(w,"更新失败",500)
        return 
    }
    w.Header().Add("Location","/update?id="+id)
    w.WriteHeader(302)
}

func DeleteController(w http.ResponseWriter,r *http.Request){
     r.ParseForm()
     id:=r.Form.Get("id")
     id_num,err:=strconv.Atoi(id)
     if err!=nil{
        http.Error(w,"参数错误",404)
        return 
    }
    err=model.DeleteById(int64(id_num))
    if err !=nil{
        http.Error(w,"删除失败",500)
    }
    w.Header().Add("Location","/")
    w.WriteHeader(302)
}

