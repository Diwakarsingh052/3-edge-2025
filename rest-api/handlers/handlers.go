package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rest-api/middleware"
)

func API() http.Handler {

	// standard lib
	//mux := http.NewServeMux()
	//return mux

	r := mux.NewRouter()
	r.Use(middleware.Logger)
	r.HandleFunc("/check", Check)

	// we can return gorilla mux router as http.Handler because it implements the type
	return r

}

func Check(w http.ResponseWriter, r *http.Request) {
	// setting the header for json content type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "ok",
	})
}
