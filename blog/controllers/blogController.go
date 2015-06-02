package controller

import (
    "net/http"
    "github.com/silenceper/go-blog/blog/models/"
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
