package counter

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Message is a data chunk from counter
type Message struct {
	Counter int
}

// #region IMessage

// ReportToLog sends message to standard log system
func (m Message) ReportToLog() {
	log.WithFields(log.Fields{
		"counter": m.Counter,
	}).Info("Received")
}

// ToString returns string representations
func (m Message) ToString() string {
	return fmt.Sprintf("Counter is %d", m.Counter)
}

// #endregion
