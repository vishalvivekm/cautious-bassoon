package main

import (
	"log"
	"net/http"
	"vishalvivekm/location-tracker/handler"
	"vishalvivekm/location-tracker/models"
)
func main() {
	h := models.NewWebSocketHub()

	http.Handle("/push", handler.PushLocation(h))
	http.Handle("/track", handler.ServeWs(h))

	log.Println("Listening at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("error starting server: %s", err)
	}


}
