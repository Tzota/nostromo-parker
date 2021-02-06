package harvester

import "github.com/go-redis/redis/v8"

// IHarvester is common interface for a sensor
type IHarvester interface {
	ListenTo(chan []byte) // wrap with a structure?
	GetMessagesChannel() chan IMessage
}

// IMessage is a data message from sensor, visitor pattern
type IMessage interface {
	// ReportToRedisStream writes data to Redis Stream
	ReportToRedisStream(*redis.Client, string) error
	// ReportToLog sends message to standard log system
	ReportToLog()
	// ToString returns string representations
	ToString() string
}
