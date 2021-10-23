package domains

// Event type
type Event struct {
	WebsiteUrl			string			`json:"websiteUrl"`
	SessionId			string 			`json:"sessionId"`
	FormCompletionTime  int  			`json:"formCompletionTime"`
	CopyAndPaste       	map[string]bool `json:"longitude_deg"`
	ResizeFrom			Dimension  		`json:"resizeFrom"`
	ResizeTo			Dimension  		`json:"resizeTo"`
}

type Dimension struct {
	Width  string	`json:"width"`
	Height string	`json:"height"`
}

