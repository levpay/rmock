package main

import (
	"io/ioutil"
	"log"
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
	return ":" + p
}

func HandleError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "some err", http.StatusBadRequest)
}

func HandleOK(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(bodyBytes))
	w.WriteHeader(http.StatusOK)
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Levpay rmock project"))
}
