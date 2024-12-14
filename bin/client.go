package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Log = map[string]any

var (
	LEVELS   []string = []string{"ERROR", "INFO", "DEBUG", "CRITICAL", "WARNING"}
	MESSAGES []string = []string{
		"Failed to generate message",
		"Started Service",
		"Created new document",
		"Failed to complete request",
		"Field is required but missing",
	}
)

func generateLog() Log {
	s := rand.NewSource(time.Hour.Milliseconds())
	r := rand.New(s)

	msgId := r.Intn(len(MESSAGES))
	levelId := r.Intn(len(LEVELS))
	reqId := r.Intn(1e6)

	return Log{
		"time":      time.Now().String(),
		"message":   MESSAGES[msgId],
		"level":     LEVELS[levelId],
		"requestId": reqId,
	}
}

func sendLog() {
	log.Printf("Sending a Log\n")
	newLog := generateLog()
	encodedData, err := json.Marshal(newLog)
	if err != nil {
		log.Printf("%s\n", err.Error())
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:6588/ingest/json", bytes.NewBuffer(encodedData))
	if err != nil {
		log.Printf("[Error] %s\n", err.Error())
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("[Error] %s\n", err.Error())
		return
	}

	if res.StatusCode != 201 {
		log.Printf("[Error] Response Code: %d\n", res.StatusCode)
		return
	}
}

func main() {
	wg := sync.WaitGroup{}

	for range 20 {
		go func() {
			wg.Add(1)
			defer wg.Done()

			sendLog()
		}()
	}

	time.Sleep(2 * time.Second)
	wg.Wait()
}
