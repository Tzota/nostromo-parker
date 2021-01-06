package ds18b20

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tzota/nostromo-parker/internal/domain"
	"github.com/tzota/nostromo-parker/internal/harvester"
	"github.com/tzota/nostromo-parker/internal/sensors"
)

const sensorDelimiter = "\r\n"

var sensorType = domain.SensorType{
	Name:        "Ds18b20",
	Description: "DS18B20 temperature sensor",
}

// Sensor is a buffered storage for data from sensor
type Sensor struct {
	sensors.AbstractSensor
	bag []byte
}

// New is .ctor
func New() Sensor {
	return Sensor{
		bag: make([]byte, 0),
		AbstractSensor: sensors.AbstractSensor{
			Messages: make(chan harvester.IMessage),
			Type:     sensorType,
		},
	}
}

// eat appends chunk of data to storage and maybe send a message to channel
func (s Sensor) eat(chunk []byte) error {
	s.bag = append(s.bag, chunk...)
	str := string(s.bag)

	if pos := strings.Index(str, sensorDelimiter); pos > -1 {
		to := (pos + len(sensorDelimiter))
		part := str[0:to]
		data, err := parseString(part)
		s.bag = s.bag[to:]
		if err != nil {
			// probably partial read
			// TODO learn about error handling
			return err
		}

		go func() {
			message := newMessage(data, s.Type)
			log.WithFields(log.Fields{"temperature": message.Temperature, "sensor type": s.Type}).Info("Sending")
			s.Messages <- message
		}()
	}

	return nil
}

// #region IHarvester

// ListenTo attaches to byte channel with serial data
func (s Sensor) ListenTo(dp chan []byte) {
	for {
		log.Trace("Receiving chunk")
		chunk := <-dp
		log.WithField("len", len(chunk)).Trace("Received chunk")
		err := s.eat(chunk)
		if err != nil {
			log.Error(err)
			// TODO once is OK (connected in the middle of the pack), twice is problem
			// count errors in Parser struct private field
			// successful read should reset error counter
			continue
		}
	}
}

// GetMessagesChannel is a getter for Messages
func (s Sensor) GetMessagesChannel() chan harvester.IMessage {
	return s.Messages
}

// #endregion
