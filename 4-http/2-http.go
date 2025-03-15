package main

import "net/http"

func main() {
	http.HandleFunc("/go", goroutine)
	http.HandleFunc("/json", sendJson)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func goroutine(w http.ResponseWriter, r *http.Request) {
	//panic("this handler func is panicking")
	go func() {
		panic("panic in the new goroutine")
	}()

}

func sendJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "hello world"}`))
}
