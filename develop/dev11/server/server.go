package server

import (
	"encoding/json"
	"fmt"
	"log"
	"mbatimel/WB_Tech-L2/tree/main/develop/dev11/tools"
	"mbatimel/WB_Tech-L2/tree/main/develop/dev11/event"
	"net/http"
	"os"
	"sync"
)

type Port struct {
	Port string `json:"port"`
}

type Server struct {
	Mu    sync.RWMutex
	Cache map[string]event.Event
	port  Port
}

func NewServer() (*Server, error) {
	data, err := os.ReadFile("config/config.json")
	if err != nil {
		return nil, err
	}
	port := Port{}
	err = json.Unmarshal(data, &port)
	if err != nil {
		return nil, err
	}
	return &Server{
		Cache: make(map[string]event.Event),
		port:  port,
	}, nil
}

func (s *Server) SetEvent(event event.Event) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Cache[event.EventName] = event
}

func (s *Server) DeleteEventsv(eventName string) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	delete(s.Cache, eventName)
}

func (s *Server) SetupGetHandlers() {
	http.HandleFunc("/event_by_name", tools.MiddlewareLogger(s.EventByName))
	http.HandleFunc("/events_for_day", tools.MiddlewareLogger(s.EventsForDay))
	http.HandleFunc("/events_for_week", tools.MiddlewareLogger(s.EventsForWeek))
	http.HandleFunc("/events_for_month", tools.MiddlewareLogger(s.EventsForMonth))
}

func (s *Server) SetupPostHandlers() {
	http.HandleFunc("/create_event", tools.MiddlewareLogger(s.CreateEvent))
	http.HandleFunc("/update_event", tools.MiddlewareLogger(s.UpdateEvent))
	http.HandleFunc("/delete_event", tools.MiddlewareLogger(s.DeleteEvent))
}

func (s *Server) SetupHandlers() {
	s.SetupGetHandlers()
	s.SetupPostHandlers()
}

func (s *Port) getAddress() string {
	return fmt.Sprintf(":%s", s.Port)
}

func (s *Server) Up() {
	address := s.port.getAddress()
	fmt.Println("Server listen on", address)
	log.Println(http.ListenAndServe(address, nil))
}