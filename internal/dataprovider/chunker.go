package dataprovider

import (
	"io"

	log "github.com/sirupsen/logrus"
)

// GetChunker reads complete chunk of data from socket
func GetChunker(conn io.Reader) chan []byte {
	c := make(chan []byte)

	go func() {
		for {
			buffer := make([]byte, 50)
			n, err := conn.Read(buffer)
			if err != nil {
				log.WithField("error", err).Error("Read from connection")
				continue
			}

			if n > 0 {
				c <- buffer[:n]
			}
		}
	}()

	return c
}
