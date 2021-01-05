package harvester

// IHarvester is common interface for a sensor
type IHarvester interface {
	ListenTo(chan []byte) // wrap with a structure?
	GetMessagesChannel() chan IMessage
}

// IMessage is a data message from sensor, visitor pattern
type IMessage interface {
	// ReportToLog sends message to standard log system
	ReportToLog()
	// ToString returns string representations
	ToString() string
}
