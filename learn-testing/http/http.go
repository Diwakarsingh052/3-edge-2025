package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func doubleHandler(w http.ResponseWriter, r *http.Request) {

	text := r.URL.Query().Get("v") // trying to fetch value from the data
	if text == "" {
		http.Error(w, "missing value", http.StatusBadRequest)
		log.Println("missing value")
		return // don't forget the return
	}

	v, err := strconv.Atoi(text)
	if err != nil {
		log.Println(err)
		http.Error(w, "not a number: "+text, http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, v*2)

	// the return values must be tested // here we have 3

}
