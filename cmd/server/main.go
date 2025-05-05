package main

import (
	"io"
	"log"
	"minisentry/configs"
	"minisentry/internal/kafka"
	"net/http"
)

func logHandler(w http.ResponseWriter, r *http.Request) {
	conf := configs.LoadConfig()
	k := kafka.NewKafkaProducer(&conf.Kafka)

	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err := k.WriteMessage(body); err != nil {
		log.Println("failed to send to Kafka:", err)
		http.Error(w, "kafka error", 500)
		return
	}
}

func main() {
	http.HandleFunc("POST /log", logHandler)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
