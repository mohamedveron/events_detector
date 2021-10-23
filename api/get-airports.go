package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) GetAirports(w http.ResponseWriter, r *http.Request) {

	fmt.Println("server handler...")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to read body"))
		return
	}

	req := &helloWorldRequest{}

	if err = json.Unmarshal(body, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to unmarshal JSON request"))
		return
	}

	log.Printf("Request received %+v", req)

	w.WriteHeader(http.StatusOK)
}
