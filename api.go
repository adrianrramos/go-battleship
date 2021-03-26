package main

import (
    "errors"
    "log"
    "net/http"
    "strings"
    "battleship/board"

    "github.com/google/uuid"
)

type Game struct {
    SessionId string `json:"session_id"`
    GameBoard board.Board `json:"game_board"`
}

var games = make(map[string]*Game)

type RegisterShot struct {
    X int `json:"x"`
    Y int `json:"y"`
    SessionId string `json:"session_id"`
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
    sessionId := strings.TrimPrefix(r.URL.Path, "/game/")
    
    switch {
        case r.Method == "GET" && sessionId != "":
            body := games[sessionId]
            encodeJSONResponse(w, body)    
            return
        case r.Method == "GET" && sessionId == "":
            encodeJSONResponse(w, games)
            return
        case r.Method == "POST":
            var p RegisterShot
            
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

            game := games[p.SessionId]
            board.RegisterShot(p.X, p.Y, &game.GameBoard)

            encodeJSONResponse(w, game)
            return
        default:
            return
    }

}

func createGame() string {
    sessionId := uuid.New().String()[:7]

    gameBoard := board.NewBoard()
    board.PlaceShips(&gameBoard)

    games[sessionId] = &Game{ sessionId, gameBoard }

    return sessionId
}
