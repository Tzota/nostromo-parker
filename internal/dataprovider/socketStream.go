package dataprovider

// UnixReader decouples source of data for testing...
type UnixReader interface {
	Read(fd int, p []byte) (n int, err error)
}

// GetChunker reads complete chunk of data from socket
func GetChunker(ur UnixReader, fd int) chan []byte {
	c := make(chan []byte)

	go func() {
		for {
			buffer := make([]byte, 50)
			n, _ := ur.Read(fd, buffer)

			if n > 0 {
				c <- buffer[:]
			}
		}
	}()

	return c
}
