package nest

// Connection contains important connection info
type Connection struct {
	AccessToken string
	testURL     string
}

const BaseURL = "https://developer-api.nest.com/devices"
