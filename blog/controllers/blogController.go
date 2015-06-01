package controller

import (
    "net/http"
    "io"
    "../models/"
    "fmt"
)

func IndexController(w http.ResponseWriter,r *http.Request){
    io.WriteString(w,"index Controller") 
    list:=model.GetList()
    fmt.Println(list)
}
