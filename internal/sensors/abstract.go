package sensors

import (
	"github.com/tzota/nostromo-parker/internal/domain"
	"github.com/tzota/nostromo-parker/internal/harvester"
)

// AbstractSensor represents one sensor
type AbstractSensor struct {
	Messages chan harvester.IMessage
	Type     domain.SensorType
}
