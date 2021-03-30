//TODO describe the response for all routes [header, status, etc] ==> google why this improves performance
package main

import (
	"battleship/board"
	"errors"
	"log"
	"net/http"
	"strings"
//    "fmt"

	"github.com/google/uuid"
)

type Game struct {
	SessionId string        `json:"session_id"`
	GameBoard *board.Board   `json:"game_board"`
}

var games = make(map[string]*Game)

type ShotData struct {
	X         int    `json:"x"`
	Y         int    `json:"y"`
	SessionId string `json:"session_id"`
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := strings.TrimPrefix(r.URL.Path, "/game/")

	switch r.Method {
	case "GET":
		if sessionId != "" {
			body := games[sessionId]
			encodeJSONResponse(w, body)
            return
		}
		encodeJSONResponse(w, games)
        return
	case "POST":
		var p ShotData

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
		game.GameBoard.RegisterShot(p.X, p.Y)

		encodeJSONResponse(w, game)
		return
	default:
		return
	}

}

func createGame() string {
	sessionId := uuid.New().String()[:7]

	gameBoard := board.NewBoard()
	gameBoard.PlaceShips()

    games[sessionId] = &Game{sessionId, gameBoard}

	return sessionId
}
