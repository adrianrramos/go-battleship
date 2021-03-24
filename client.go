package main

import (
//    "fmt"
    "log"
    "html/template"
    "net/http"
    "path"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Incoming request from: ", r.RemoteAddr)
    fp := path.Join("templates", "index.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                        return
                            
    }

    if err := tmpl.Execute(w, ""); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                    
    }
}
