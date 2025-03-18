package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func API() http.Handler {

	// standard lib
	//mux := http.NewServeMux()
	//return mux

	r := mux.NewRouter()
	r.HandleFunc("/check", Check)

	// we can return gorilla mux router as http.Handler because it implements the type
	return r

}

func Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "ok",
	})
}
