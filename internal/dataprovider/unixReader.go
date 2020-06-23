package dataprovider

import "golang.org/x/sys/unix"

// RealUnixReader это настоящая имплементация UnixReader
type RealUnixReader struct{}

func (r RealUnixReader) Read(fd int, p []byte) (n int, err error) {
	return unix.Read(fd, p)
}
