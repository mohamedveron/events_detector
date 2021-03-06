package service

import (
	"github.com/mohamedveron/events_detector/domains"
	"log"
)

func (s *Service)  HandleCopyAndPasteEvent(e domains.EventRequest) (domains.Event,error){

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

	log.Printf("Copy and paste Event triggered %+v", s.Events[e.SessionId])

	return s.Events[e.SessionId], nil
}

func (s *Service)  HandleFormSubmissionEvent(e domains.EventRequest) (domains.Event,error){

	copyAndPastMap := make(map[string]bool)

	// check if the struct is already constructed
	if _, ok := s.Events[e.SessionId]; ok {

		copyAndPastMap = s.Events[e.SessionId].CopyAndPaste

		s.mutex.Lock()

		event := s.Events[e.SessionId]
		event.FormCompletionTime = e.TimeTaken
		s.Events[e.SessionId] = event
		s.mutex.Unlock()

	}else{

		event := domains.Event{
			WebsiteUrl:         e.WebsiteUrl,
			SessionId:          e.SessionId,
			EventType:          e.EventType,
			TimeTaken:          0,
			FormCompletionTime: e.TimeTaken,
			CopyAndPaste:       copyAndPastMap,
			ResizeFrom:         domains.Dimension{},
			ResizeTo:           domains.Dimension{},
		}

		s.mutex.Lock()
		s.Events[event.SessionId] = event
		s.mutex.Unlock()

	}

	log.Printf("Form Submission Event triggered %+v", s.Events[e.SessionId])

	return s.Events[e.SessionId], nil
}

func (s *Service)  HandleScreenResizeEvent(e domains.EventRequest) (domains.Event,error){

	// check if the struct is already constructed
	if _, ok := s.Events[e.SessionId]; ok {

		s.mutex.Lock()

		event := s.Events[e.SessionId]
		event.ResizeFrom = e.ResizeFrom
		event.ResizeTo 	= e.ResizeTo
		s.Events[e.SessionId] = event

		s.mutex.Unlock()
	}else{

		copyAndPastMap := make(map[string]bool)

		event := domains.Event{
			WebsiteUrl:         e.WebsiteUrl,
			SessionId:          e.SessionId,
			EventType:          e.EventType,
			TimeTaken:          0,
			FormCompletionTime: 0,
			CopyAndPaste:       copyAndPastMap,
			ResizeFrom:         e.ResizeFrom,
			ResizeTo:           e.ResizeTo,
		}

		s.mutex.Lock()
		s.Events[event.SessionId] = event
		s.mutex.Unlock()

	}

	log.Printf("Screen size Event triggered %+v", s.Events[e.SessionId])

	return s.Events[e.SessionId], nil
}
