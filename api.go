package main

import (
    "errors"
    "log"
    "encoding/json"
    "net/http"
)

type PostBody  struct {
    Message string
    Weight int
    Title string
}

func test(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case "GET":
            log.Println("Incoming request from", r.RemoteAddr)
            w.Header().Set("Content-Type", "application/json")

            body := make(map[string]string) 
            body["message"] = "hello joe"

            json.NewEncoder(w).Encode(body)

        case "POST":
            var p PostBody
            
            err := decodeJSONBody(w, r, &p)
            if err != nil {
                var mr *malformedRequest
                if errors.As(err, &mr) {
                    http.Error(w, mr.msg, mr.status)
                } else {
                    log.Println(err.Error())
                    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
                }
                    return
            }

            log.Println(p)
        default:
            return
    }
}


