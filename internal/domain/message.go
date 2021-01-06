package domain

// AbstractMessage is a base message from sensor
type AbstractMessage struct {
	SensorType SensorType
}

// NewMessage is ctor
func NewMessage(s SensorType) AbstractMessage {
	return AbstractMessage{s}
}
