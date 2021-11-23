package dataprovider

import (
	"io"

	log "github.com/sirupsen/logrus"
)

// GetChunker reads complete chunk of data from socket
func GetChunker(conn io.Reader) (chan []byte, chan bool) {
	c := make(chan []byte)
	flag := make(chan bool)

	go func() {
		for {
			buffer := make([]byte, 50)
			n, err := conn.Read(buffer)
			if err != nil {
				log.WithField("error", err).Error("Read from connection")
				close(c)
				defer (func() {
					flag <- true
					close(flag)
				})()

				// continue
				return
			}

			if n > 0 {
				c <- buffer[:n]
			}
		}
	}()

	return c, flag
}
