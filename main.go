package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/err", HandleError)
	http.HandleFunc("/ok", HandleOK)
	http.HandleFunc("/", HandleRoot)
	http.ListenAndServe(port(), nil)
}

func port() string {
	p := os.Getenv("PORT")
	if p == "" {
		return ":3333"
	}
	return p
}

func HandleError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "some err", http.StatusBadRequest)
}

func HandleOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Levpay rmock project"))
}
