package ds18b20

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

const sensorDelimiter = "\r\n"

// Harvester is a buffered storage for data from sensor
type Harvester struct {
	bag      []byte
	Messages chan Message
}

// New is .ctor
func New() Harvester {
	return Harvester{
		bag:      make([]byte, 0),
		Messages: make(chan Message),
	}
}

// ListenTo attaches to byte channel with serial data
func (p Harvester) ListenTo(dp chan []byte) {
	for {
		log.Info("gonna receive chunk")
		chunk := <-dp
		log.Info("chunk received")
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

// eat appends chunk of data to storage and maybe send a message to channel
func (p Harvester) eat(chunk []byte) error {
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
				"temperature": message.Temperature,
			}).Infof("Sending ds18b20 message")
			p.Messages <- message
		}()
	}

	return nil
}
