package api

import (
	"github.com/mohamedveron/events_detector/service"
)

type Server struct {
	svc *service.Service
}

func NewServer(svc *service.Service) *Server {
	return &Server{
		svc: svc,
	}
}
