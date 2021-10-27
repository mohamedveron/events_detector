package domains

// Event type
type Event struct {
	WebsiteUrl			string
	SessionId			string
	EventType  			string
	TimeTaken  			int
	FormCompletionTime  int
	CopyAndPaste       	map[string]bool
	ResizeFrom			Dimension
	ResizeTo			Dimension
}

type Dimension struct {
	Width  string	`json:"width"`
	Height string	`json:"height"`
}


type EventRequest struct {
	WebsiteUrl			string			`json:"websiteUrl"`
	SessionId			string 			`json:"sessionId"`
	EventType  			string 			`json:"eventType"`
	TimeTaken  			int 			`json:"timeTaken"`
	FormId  			string  		`json:"formId"`
	ResizeFrom			Dimension  		`json:"resizeFrom"`
	ResizeTo			Dimension  		`json:"resizeTo"`
	Pasted				bool			`json:"pasted"`
}
