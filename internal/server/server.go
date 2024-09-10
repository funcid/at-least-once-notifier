package server

import (
	"encoding/json"
	"log"
	"net/http"

	_ "at-least-once-notifier/internal/model"
	"at-least-once-notifier/internal/notifier"

	"github.com/gorilla/mux"
)

type Server struct {
	notifyService *notifier.NotificationService
}

func NewServer(notifyService *notifier.NotificationService) *Server {
	return &Server{notifyService: notifyService}
}

func (s *Server) Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/notify", s.handleNotify).Methods("POST")
	return router
}

func (s *Server) handleNotify(w http.ResponseWriter, r *http.Request) {
	var outboxEntry notifier.OutboxEntry
	err := json.NewDecoder(r.Body).Decode(&outboxEntry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.notifyService.AddToOutbox(outboxEntry)
	if err != nil {
		log.Printf("Error adding to outbox: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notification added to outbox"))
}
