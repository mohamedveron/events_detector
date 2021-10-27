package tests

import (
	"github.com/mohamedveron/events_detector/domains"
	"github.com/mohamedveron/events_detector/service"
	"sync"
	"testing"
)


func TestTriggerCopyPasteEvent(t *testing.T) {

	eventRequest := domains.EventRequest{
		WebsiteUrl: "https://www.ravelin.com/",
		SessionId:  "5553wrwq5tgbvfr4",
		EventType:  "copyAndPaste",
		TimeTaken:  0,
		FormId:     "inputCardNumber",
		ResizeFrom: domains.Dimension{},
		ResizeTo:   domains.Dimension{},
		Pasted:     true,
	}

	serviceLayer := service.NewService()

	eventResponse, err := serviceLayer.HandleCopyAndPasteEvent(eventRequest)

	if err != nil {
		t.Errorf("Can't trigger the event")
	}

	assertStringEquals(t, eventRequest.SessionId, eventResponse.SessionId)

}

func TestTriggerTimeTakenEvent(t *testing.T) {

	eventRequest := domains.EventRequest{
		WebsiteUrl: "https://www.ravelin.com/",
		SessionId:  "5553wrwq5tgbvfr4",
		EventType:  "timeTaken",
		TimeTaken:  54,
		FormId:     "inputCardNumber",
		ResizeFrom: domains.Dimension{},
		ResizeTo:   domains.Dimension{},
		Pasted:     false,
	}

	serviceLayer := service.NewService()

	eventResponse, err := serviceLayer.HandleFormSubmissionEvent(eventRequest)

	if err != nil {
		t.Errorf("Can't trigger the event")
	}

	assertStringEquals(t, eventRequest.SessionId, eventResponse.SessionId)

}

func TestTriggerScreenResizeEvent(t *testing.T) {

	eventRequest := domains.EventRequest{
		WebsiteUrl: "https://www.ravelin.com/",
		SessionId:  "5553wrwq5tgbvfr4",
		EventType:  "timeTaken",
		TimeTaken:  0,
		FormId:     "inputCardNumber",
		ResizeFrom: domains.Dimension{
			Width: "234",
			Height: "165",
		},
		ResizeTo:   domains.Dimension{
			Width: "300",
			Height: "200",
		},
		Pasted:     false,
	}

	serviceLayer := service.NewService()

	eventResponse, err := serviceLayer.HandleScreenResizeEvent(eventRequest)

	if err != nil {
		t.Errorf("Can't trigger the event")
	}

	assertStringEquals(t, eventRequest.SessionId, eventResponse.SessionId)

}

func TestTriggerCopyPasteEventConcurrently(t *testing.T) {

	eventRequest := domains.EventRequest{
		WebsiteUrl: "https://www.ravelin.com/",
		SessionId:  "5553wrwq5tgbvfr4",
		EventType:  "copyAndPaste",
		TimeTaken:  0,
		FormId:     "inputCardNumber",
		ResizeFrom: domains.Dimension{},
		ResizeTo:   domains.Dimension{},
		Pasted:     true,
	}

	eventRequest2 := domains.EventRequest{
		WebsiteUrl: "https://www.ravelin.com/",
		SessionId:  "5553wrwq5tgbvfr4",
		EventType:  "copyAndPaste",
		TimeTaken:  0,
		FormId:     "inputCVV",
		ResizeFrom: domains.Dimension{},
		ResizeTo:   domains.Dimension{},
		Pasted:     true,
	}

	serviceLayer := service.NewService()

	var wg sync.WaitGroup
	wg.Add(3)

	go serviceLayer.HandleCopyAndPasteEvent(eventRequest)
	go serviceLayer.HandleCopyAndPasteEvent(eventRequest2)

	wg.Wait()

	eventRequest.TimeTaken = 180
	eventResponse, err := serviceLayer.HandleFormSubmissionEvent(eventRequest)

	if err != nil {
		t.Errorf("Can't trigger the event")
	}

	assertStringEquals(t, eventRequest.SessionId, eventResponse.SessionId)

}

func assertStringEquals(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Errorf("Actual: %s, Expected: %s", actual, expected)
	}
}
