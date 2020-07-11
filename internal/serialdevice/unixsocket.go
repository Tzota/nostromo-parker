package serialdevice

import (
	"syscall"

	"golang.org/x/sys/unix"
)

// UnixSocket is a real world unix implementation
type UnixSocket struct {
}

// Socket creates unix socket
func (u UnixSocket) Socket() (int, error) {
	return unix.Socket(syscall.AF_BLUETOOTH, syscall.SOCK_STREAM, unix.BTPROTO_RFCOMM)
}

// Connect creates connection between socket and address
func (u UnixSocket) Connect(fd int, addr *unix.SockaddrRFCOMM) error {
	return unix.Connect(fd, addr)
}

// Close closes socket
func (u UnixSocket) Close(fd int) error {
	return unix.Close(fd)
}

// Read reads socket
func (u UnixSocket) Read(fd int, p []byte) (n int, err error) {
	return unix.Read(fd, p)
}
