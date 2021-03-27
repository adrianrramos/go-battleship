package main

import (
	"net/http"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate("index", w, "")
}

// TODO: "game/new/" will redirect to cached session, therefore not actually generating a unique session
func sessionHandler(w http.ResponseWriter, r *http.Request) {
	session_id := strings.TrimPrefix(r.URL.Path, "/game/")
	if session_id == "new" {
		// generate new game
		session_id = createGame()
		http.Redirect(w, r, "/game/"+session_id, 301)
	}

	renderTemplate("game", w, session_id)
}
