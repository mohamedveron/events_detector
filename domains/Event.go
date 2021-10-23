package domains

// Event type
type Event struct {
	WebsiteUrl string `json:"websiteUrl"`
	SessionId       string `json:"sessionId"`
	FormCompletionTime   int  `json:"formCompletionTime"`
	Longitude   float64  `json:"longitude_deg"`
	Address string
	CopyAndPaste       map[string]bool `json:"longitude_deg"`
}

type Data struct {
	WebsiteUrl         string
	SessionId          string
	ResizeFrom         Dimension
	ResizeTo           Dimension
	CopyAndPaste       map[string]bool // map[fieldId]true
	FormCompletionTime int // Seconds
}

type Dimension struct {
	Width  string
	Height string
}

