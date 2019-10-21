package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
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
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Fprintf(w, "userip: %q is not IP:port", r.RemoteAddr)
	}
	userIP := net.ParseIP(ip)
	log.Printf("%+v\n", userIP)
	log.Println(string(bodyBytes))
	w.WriteHeader(http.StatusOK)
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Levpay rmock project"))
}
