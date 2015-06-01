package blog

import (
    "net/http"
    "./controllers/"
    "./config/"
)

func parseConfig(){

}

func Run(){
    //parse config.json 

    //set static dir  
    http.Handle("/static/",http.FileServer(http.Dir(config.CFG_STATIC_DIR)))
    
    http.HandleFunc("/",controller.IndexController)
    
    http.ListenAndServe(":9000",nil)
}
