package service

import (
	"github.com/mohamedveron/events_detector/domains"
	"sync"
)

type Service struct {
	Events	map[string]domains.Event
	mutex sync.Mutex
}

func NewService() *Service {

	events := make(map[string]domains.Event)

	return &Service{
		Events: events,
	}
}
