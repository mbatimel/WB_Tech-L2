package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"mbatimel/WB_Tech-L2/tree/main/develop/dev11/event"
	"mbatimel/WB_Tech-L2/tree/main/develop/dev11/tools"
)

const (
	permissionError     int = 2
	valid               int = 1
	invalidData         int = 0
	internalServerError int = -1
)

func GetDataFromRequest(r *http.Request) (event.Event, error) {
	event := event.Event{}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return event, err
	}
	err = json.Unmarshal(data, &event)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (s *Server) ValidatePost(w http.ResponseWriter, event *event.Event, actionType string) int {
	s.Mu.RLock()
	data, ok := s.Cache[event.EventName]
	s.Mu.RUnlock()
	result := invalidData

	date := tools.IsValidDate(event)
	switch actionType {
	case "create":
		if !ok && date && event.EventName != "" {
			result = valid
		}
	case "update":
		if ok && date {
			result = valid
		}
	case "delete":
		if ok {
			if data.UserId == event.UserId {
				result = valid
			} else {
				result = permissionError
			}
		}
	default:
		result = internalServerError
	}
	return result
}

func (s *Server) ValidateAndRespond(w http.ResponseWriter, code int) bool {
	if code == valid {
		return true
	}
	switch code {
	case internalServerError:
		tools.MakeJsonRespond(w, 503, tools.JsonError("internal server error"))
	case invalidData:
		tools.MakeJsonRespond(w, 400, tools.JsonError("invalid data"))
	case permissionError:
		tools.MakeJsonRespond(w, 500, tools.JsonError("permisson error"))
	}
	return false
}

func (s *Server) PostRequestCheck(w http.ResponseWriter, r *http.Request, request string) (event.Event, error) {
	event :=event.Event{}
	if r.Method != http.MethodPost {
		errorString := "method not allowed"
		tools.MakeJsonRespond(w, 500,tools. JsonError(errorString))
		return event, fmt.Errorf(errorString)
	}
	event, err := GetDataFromRequest(r)
	if err != nil {
		log.Println(err)
		tools.MakeJsonRespond(w, 503, tools.JsonError("internal server error"))
		return event, err
	}
	validate := s.ValidatePost(w, &event, request)
	if !s.ValidateAndRespond(w, validate) {
		return event, fmt.Errorf("something being wrong")
	}
	return event, nil
}

func (s *Server) CreateAndUpdate(w http.ResponseWriter, r *http.Request, request string) {
	event, err := s.PostRequestCheck(w, r, request)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.SetEvent(event)
	tools.MakeJsonRespond(w, 200, tools.JsonResult("ok"))
}

func (s *Server) CreateEvent(w http.ResponseWriter, r *http.Request) {
	s.CreateAndUpdate(w, r, "create")
}

func (s *Server) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	s.CreateAndUpdate(w, r, "update")
}

func (s *Server) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	event, err := s.PostRequestCheck(w, r, "delete")
	if err != nil {
		return
	}
	s.DeleteEventsv(event.EventName)
	tools.MakeJsonRespond(w, 200, tools.JsonResult("ok"))
}