package main

import (
	"log"
	"net/http"
	"io"
	"time"
)

var version string

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func ip(client *http.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := client.Get("https://ip.mrfriday.com")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer res.Body.Close()
		b, err := io.ReadAll(res.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(b)
	}
}

func main() {
	log.SetFlags(0)
	client := &http.Client{Timeout: time.Second * 5}
	http.HandleFunc("/health", health)
	http.HandleFunc("/ip", ip(client))
	log.Printf("server version %s started", version)
	log.Printf("listening on port 8989")
	log.Fatal(http.ListenAndServe(":8989", nil))
}
