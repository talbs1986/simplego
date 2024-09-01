package messaging

// MessageType - defines the type of message
type MessageType string

// MessageWrapper represents a single message payload wrapper object
type MessageWrapper struct {
	// id - unique message id
	ID string `json:"id"`
	// sender - unique sender name
	Sender string `json:"sender"`
	// ts - message published timestamp milliseconds
	Timestamp int64 `json:"ts"`
	// eventType - event type
	EventType MessageType `json:"eventType"`
	// payload - message payload
	Payload interface{} `json:"payload"`
}
