package controller

import (
    "path"
    "bytes"
    "github.com/silenceper/go-blog/blog/config"
    "html/template"
)

var FuncMap template.FuncMap

func init(){
    FuncMap=make(template.FuncMap)
}

func Render(tpl string,data map[string]interface{}) ([]byte,error){
    var (
        tpl_path=path.Join(config.CFG_TPL_DIR,tpl)
        t *template.Template
        err error
    )
    t = template.New(path.Base(tpl_path)).Funcs(FuncMap)
    t,err=t.ParseFiles(tpl_path)
    if err!=nil{
        return nil,err
    }
    var buf bytes.Buffer
    err=t.Execute(&buf,data) 
    if err!=nil{
        return nil,err
    }
    return buf.Bytes(),nil
}
