package main 

import (
    "fmt"
    "net/http"
    "log"
    "sync"
)

func main () {
    wg := new(sync.WaitGroup)
    wg.Add(2)

    apiServer := http.NewServeMux()
    apiServer.HandleFunc("/test", test)
    
    clientServer := http.NewServeMux()
    clientServer.HandleFunc("/", indexHandler)

    go func () {
        port := ":3001"
        fmt.Println("API server listening on port =>", port)
        log.Fatal(http.ListenAndServe(port, apiServer))
        wg.Done()
    }()

    go func () {
        port := ":3000"
        fmt.Println("Client server listening on port =>", port)
        log.Fatal(http.ListenAndServe(port, clientServer))
        wg.Done()
    }()   

    wg.Wait()
}
