package api

import (
	"encoding/json"
	"fmt"
	"github.com/mohamedveron/events_detector/domains"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) AddEvent(w http.ResponseWriter, r *http.Request) {

	fmt.Println("server handler...")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to read body"))
		return
	}

	req := &domains.Event{}

	if err = json.Unmarshal(body, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to unmarshal JSON request"))
		return
	}

	log.Printf("Request received %+v", req)

	w.WriteHeader(http.StatusOK)
}
