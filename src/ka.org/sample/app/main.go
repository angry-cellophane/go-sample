package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main()  {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("up"))
	})

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(fmt.Sprintf("hostname=%s", hostname)))
	})

	app := http.Server{Addr: ":8080"}
	if err := app.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
