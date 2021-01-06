package counter

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/tzota/nostromo-parker/internal/domain"
)

type counterPayload struct {
	Counter int
}

// Message is a data chunk from counter
type Message struct {
	domain.AbstractMessage
	counterPayload
}

func newMessage(p counterPayload, st domain.SensorType) Message {
	return Message{
		domain.NewMessage(st),
		p,
	}
}

// #region IMessage

// ReportToLog sends message to standard log system
func (m Message) ReportToLog() {
	log.WithFields(log.Fields{"counter": m.Counter, "sensor type": m.SensorType}).Info("Received")
}

// ToString returns string representations
func (m Message) ToString() string {
	return fmt.Sprintf("Counter is %d", m.Counter)
}

// #endregion
