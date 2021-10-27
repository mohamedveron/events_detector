package service

import "github.com/mohamedveron/events_detector/domains"

func (s *Service)  handleCopyAndPasteEvent(e domains.EventRequest) error{

	copyAndPastMap := make(map[string]bool)

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
	
	return nil
}
