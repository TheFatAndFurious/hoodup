package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

var messages chan string
func init() {
	messages = make(chan string, 10)
}

func EventsHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Type")
		w.Header().Set("Content-Encoding", "identity")

		defer func() {
			log.Println("closing connection")
			if recover() != nil {
				log.Println("recovered from panic, closing channel")
				close(messages)
			}		
		}()


			for {
				select {
				case msg := <-messages:
					fmt.Fprintf(w, "data: %v\n\n", msg)
					if flusher, ok := w.(http.Flusher); ok {
						flusher.Flush()
                    }
				case <-r.Context().Done():
					log.Println("context done")
					return
				}
			}

	}
}