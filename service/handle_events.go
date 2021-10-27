package service

import (
	"github.com/mohamedveron/events_detector/domains"
	"log"
)

func (s *Service)  HandleCopyAndPasteEvent(e domains.EventRequest) error{

	// check if the struct is already constructed
	if _, ok := s.Events[e.SessionId]; ok {

		s.mutex.Lock()
		s.Events[e.SessionId].CopyAndPaste[e.FormId] = true
		s.mutex.Unlock()
	}else{

		copyAndPastMap := make(map[string]bool)
		copyAndPastMap[e.FormId] = true

		event := domains.Event{
			WebsiteUrl:         e.WebsiteUrl,
			SessionId:          e.SessionId,
			EventType:          e.EventType,
			TimeTaken:          0,
			FormCompletionTime: 0,
			CopyAndPaste:       copyAndPastMap,
			ResizeFrom:         domains.Dimension{},
			ResizeTo:           domains.Dimension{},
		}

		s.mutex.Lock()
		s.Events[event.SessionId] = event
		s.mutex.Unlock()

	}

	log.Printf("Copy and paste Event triggered for sessionId %+v for input %+v", e.SessionId, e.FormId)

	return nil
}
