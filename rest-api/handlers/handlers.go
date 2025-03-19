package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"rest-api/auth"
	"rest-api/middleware"
)

func API(a *auth.Auth) (http.Handler, error) {

	// standard lib
	//mux := http.NewServeMux()
	//return mux

	r := mux.NewRouter()
	m, err := middleware.NewMid(a)
	if err != nil {
		return nil, err
	}
	r.Use(middleware.Logger)

	// public route, without path prefix, like homepage
	r.HandleFunc("/check", Check)

	userRouter := r.PathPrefix("/user").Subrouter()
	// signup doesn't need auth, it would be accessible
	userRouter.HandleFunc("/signup", signup)

	userAuthenticatedRouter := userRouter.NewRoute().Subrouter()
	//we are creating a new router so we can apply authentication to the specific routes
	userAuthenticatedRouter.Use(m.Authenticate)
	userAuthenticatedRouter.HandleFunc("/fetch/{email}", getUser)
	userAuthenticatedRouter.HandleFunc("/check", Check)

	// we can return gorilla mux router as http.Handler because it implements the type
	return r, nil

}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "get user ")
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "signup user ")
}

func Check(w http.ResponseWriter, r *http.Request) {
	// setting the header for json content type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "ok",
	})
}
