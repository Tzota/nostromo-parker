package counter

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tzota/nostromo-parker/internal/harvester"
)

const sensorDelimiter = "\r\n"

// Sensor is a buffered storage for data from sensor
type Sensor struct {
	bag      []byte
	Messages chan harvester.IMessage
}

// New is .ctor
func New() Sensor {
	return Sensor{
		bag:      make([]byte, 0),
		Messages: make(chan harvester.IMessage),
	}
}

// eat appends chunk of data to storage and maybe send a message to channel
func (p Sensor) eat(chunk []byte) error {
	p.bag = append(p.bag, chunk...)
	str := string(p.bag)

	if pos := strings.Index(str, sensorDelimiter); pos > -1 {
		to := (pos + len(sensorDelimiter))
		part := str[0:to]
		message, err := parseString(part)
		p.bag = p.bag[to:]
		if err != nil {
			// probably partial read
			// TODO learn about error handling
			return err
		}

		go func() {
			log.WithFields(log.Fields{
				"counter": message.Counter,
			}).Infof("Sending counter message")
			p.Messages <- message
		}()
	}

	return nil
}

// #region IHarvester

// ListenTo attaches to byte channel with serial data
func (p Sensor) ListenTo(dp chan []byte) {
	for {
		log.Trace("Receiving chunk")
		chunk := <-dp
		log.WithField("len", len(chunk)).Trace("Received chunk")
		err := p.eat(chunk)
		if err != nil {
			log.Error(err)
			// TODO once is OK (connected in the middle of the packed), twice is problem
			// count errors in Parser struct private field
			// successful read should reset error counter
			continue
		}
	}
}

// GetMessagesChannel is a getter for Messages
func (p Sensor) GetMessagesChannel() chan harvester.IMessage {
	return p.Messages
}

// #endregion
