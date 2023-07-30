package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func eventHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	defer func() {
		if r := recover(); r != nil {
			log.Println("client disconnected")
		}
	}()

	for {
		_, err := fmt.Fprintf(w, "event: message\n")
		if err != nil {
			panic(err)
		}

		_, err = fmt.Fprintf(w, "data: the time is %v\n\n", time.Now().UTC())
		if err != nil {
			panic(err)
		}

		flusher.Flush()

		time.Sleep(time.Second)
	}
}

func main() {
	http.HandleFunc("/events", eventHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
