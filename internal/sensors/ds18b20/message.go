package ds18b20

import log "github.com/sirupsen/logrus"

// Message is a data chunk from ds18b20
type Message struct {
	Temperature float32
}

// #region IMessage

// ReportToLog sends message to standard log system
func (m Message) ReportToLog() {
	log.WithFields(log.Fields{
		"temperature": m.Temperature,
	}).Info("Received")
}

// #endregion
