package ds18b20

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
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

func (m Message) toMap() map[string]interface{} {
	return map[string]interface{}{"temperature": m.Temperature, "sensor type": m.SensorType.Name}
}

// #region IMessage

// ReportToRedisStream writes data to Redis Stream
func (m Message) ReportToRedisStream(rdb *redis.Client, stream string) error {
	res, err := rdb.XAdd(context.Background(), &redis.XAddArgs{Stream: stream, MaxLenApprox: 100, ID: "*", Values: m.toMap()}).Result()
	log.WithField("res", res).Info("XADD")
	return err
}

// ReportToLog sends message to standard log system
func (m Message) ReportToLog() {
	log.WithFields(m.toMap()).Info("Received")
}

// ToString returns string representations
func (m Message) ToString() string {
	return fmt.Sprintf("Temperature is %f", m.Temperature)
}

// #endregion
