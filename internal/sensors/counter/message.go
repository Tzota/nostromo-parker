package counter

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
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

func (m Message) toMap() map[string]interface{} {
	return map[string]interface{}{
		"counter":     m.Counter,
		"sensor type": m.SensorType,
	}
}

// #region IMessage

// ReportToLog sends message to standard log system
func (m Message) ReportToLog() {
	log.WithFields(m.toMap()).Info("Received")
}

// ToString returns string representations
func (m Message) ToString() string {
	return fmt.Sprintf("Counter is %d", m.Counter)
}

// ReportToRedisStream writes data to Redis Stream
func (m Message) ReportToRedisStream(rdb *redis.Client, stream string) error {
	res, err := rdb.XAdd(context.Background(), &redis.XAddArgs{Stream: stream, MaxLenApprox: 10, ID: "*", Values: m.toMap()}).Result()
	log.WithField("res", res).Info("XADD")
	return err
}

// #endregion
