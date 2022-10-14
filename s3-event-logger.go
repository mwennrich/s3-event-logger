package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

func main() {
	var m sync.Mutex
	server := &http.Server{
		Addr:              ":5001",
		ReadHeaderTimeout: 300 * time.Second,
	}

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		b, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		events := events.S3Event{}
		err = json.Unmarshal(b, &events)
		if err != nil {
			log.Println(err)
			return
		}
		for _, record := range events.Records {
			j, err := json.Marshal(record)
			if err != nil {
				log.Println(err)
				return
			}
			m.Lock()
			fmt.Printf("%s\n", j)
			m.Unlock()
		}
	}))
	log.Fatal(server.ListenAndServe())
}
