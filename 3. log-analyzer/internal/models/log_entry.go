package models

// LogEntry represents a single parsed HTTP log entry with all relevant fields extracted
type LogEntry struct {
	PodName          string
	ContainerName    string
	IPAddress        string
	Timestamp        string
	RequestMethod    string
	RequestPath      string
	HTTPProtocol     string
	HTTPResponseCode string
	BytesSent        int
	Referer          string
	ClientInfo       string
	ServiceIP        string
	TraceID          string
}
