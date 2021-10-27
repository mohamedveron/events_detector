package tests

import (
	"github.com/mohamedveron/events_detector/domains"
	"github.com/mohamedveron/events_detector/service"
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

func assertStringNil(t *testing.T, actual string) {
	if &actual != nil {
		t.Errorf("Expected string to be nil but was %s", actual)
	}
}

func assertStringEquals(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Errorf("Actual: %s, Expected: %s", actual, expected)
	}
}
