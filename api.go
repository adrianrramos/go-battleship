package main

import (
    "errors"
    "log"
    "net/http"
    "strings"
    "battleship/board"

    "github.com/google/uuid"
)

var games map[string]string

type PostBody  struct {
    Message string
    Weight int
    Title string
}

func init() {
    games = make(map[string]string)
}

func test(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case "GET":
            body := make(map[string]string) 
            body["message"] = "hello joe"
            
            encodeJSONResponse(w, body)
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

func gameHandler(w http.ResponseWriter, r *http.Request) {
    session_id := strings.TrimPrefix(r.URL.Path, "/game/")
    
    switch {
        case session_id != "":
            body := make(map[string]string) 
            body["message"] =  games[session_id]

            encodeJSONResponse(w, body)
        case session_id == "":
            encodeJSONResponse(w, games)
        default:
            return
    }
}

func createGame() string {
    session_id := uuid.New().String()[:7]

    game_board := board.NewBoard()
    board.PlaceShips(&game_board)
    log.Println(game_board)

    games[session_id] = "This is a game with ID: " + session_id
    return session_id
}
