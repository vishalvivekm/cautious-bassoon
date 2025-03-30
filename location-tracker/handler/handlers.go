package handler

import (
	"encoding/json"
	"net"
	"net/http"
	"vishalvivekm/location-tracker/models"

	"github.com/gorilla/websocket"
)

func PushLocation(h models.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		var location models.Location
		if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
			http.Error(w, "Invalid JSON Payload", http.StatusBadRequest)
		}
		defer r.Body.Close()

		ip, _, err := net.SplitHostPort(r.RemoteAddr) // host, port, error
		if err != nil {
			http.Error(w, "Invalid client id", http.StatusBadRequest)
			return
		}

		// NOTE:
		location.DevideId = ip // ideally should come frmo jwt or session value

		if err := h.Push(location); err != nil {
			http.Error(w, "failed to push location", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Location pushed successfully"))
	}
}

func ServeWs(h models.Hub) http.HandlerFunc {

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // allow all origins
		},
	}
	return func(w http.ResponseWriter, r *http.Request){
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "failed to upgrade connection", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		clientId := r.RemoteAddr

		h.Register(clientId, conn)

		// deregister client on conenction close
		defer h.DeRegister(clientId)

		// listen for incoming messages
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				break
			}
		}
	}
}