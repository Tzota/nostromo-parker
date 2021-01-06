package ds18b20

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/tzota/nostromo-parker/internal/domain"
)

type ds18b20Payload struct {
	Temperature float32
}

// Message is a data chunk from ds18b20
type Message struct {
	domain.AbstractMessage
	ds18b20Payload
}

func newMessage(p ds18b20Payload, st domain.SensorType) Message {
	return Message{
		domain.NewMessage(st),
		p,
	}
}

// #region IMessage

// ReportToLog sends message to standard log system
func (m Message) ReportToLog() {
	log.WithFields(log.Fields{"temperature": m.Temperature, "sensor type": m.SensorType.Name}).Info("Received")
}

// ToString returns string representations
func (m Message) ToString() string {
	return fmt.Sprintf("Temperature is %f", m.Temperature)
}

// #endregion
