package kind

import (
	"fmt"

	"github.com/tzota/nostromo-parker/internal/harvester"
	"github.com/tzota/nostromo-parker/internal/sensors/counter"
	"github.com/tzota/nostromo-parker/internal/sensors/ds18b20"
)

// Kind is the enum of sensor kinds
type Kind string

const (
	// Ds18b20 temperature sensor
	Ds18b20 = "ds18b20"
	// Counter is a simple counter beacon
	Counter = "counter"
)

// New is a kinda factory
func (k Kind) New() harvester.IHarvester {
	switch k {
	case Ds18b20:
		return ds18b20.New()
	case Counter:
		return counter.New()
	default:
		panic(fmt.Sprintf("unknown kind '%s'", k))
	}
}
