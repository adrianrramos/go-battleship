package main 

import (
    "fmt"
    "net/http"
    "log"
    "sync"

    "github.com/rs/cors"
)

func main () {
    wg := new(sync.WaitGroup)
    wg.Add(2)

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000"},
        AllowCredentials: true,
        Debug: false,
    })

    apiServer := http.NewServeMux()
    apiServer.HandleFunc("/game/", gameHandler)
    go func () {
        port := ":3001"
        fmt.Println("API server listening on port =>", port)
        log.Fatal(http.ListenAndServe(port, c.Handler(apiServer)))
        wg.Done()
    }()

    clientServer := http.NewServeMux()
    clientServer.HandleFunc("/", indexHandler)
    clientServer.HandleFunc("/game/", sessionHandler)
    go func () {
        port := ":3000"
        fmt.Println("Client server listening on port =>", port)
        log.Fatal(http.ListenAndServe(port, clientServer))
        wg.Done()
    }()   

    wg.Wait()
}
