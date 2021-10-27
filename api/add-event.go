package api

import (
	"encoding/json"
	"github.com/mohamedveron/events_detector/domains"
	"io/ioutil"
	"net/http"
)

func (s *Server) AddEvent(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to read body"))
		return
	}

	req := &domains.EventRequest{}

	if err = json.Unmarshal(body, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to unmarshal JSON request"))
		return
	}

	switch eventType := req.EventType; eventType {
	case "copyAndPaste":
		s.svc.HandleCopyAndPasteEvent(*req)
	case "timeTaken":
		s.svc.HandleFormSubmissionEvent(*req)
	}

	w.WriteHeader(http.StatusOK)
}
