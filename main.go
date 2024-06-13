package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello Devlink!"))
	if err != nil {
		return
	}
}
