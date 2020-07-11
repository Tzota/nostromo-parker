package serialdevice

import (
	"errors"
	"testing"

	"golang.org/x/sys/unix"
)

type BadMac struct{}

func (u BadMac) Socket() (int, error)                            { return 0, errors.New("BadMac Socket") }
func (u BadMac) Connect(fd int, addr *unix.SockaddrRFCOMM) error { return errors.New("BadMac Connect") }
func (u BadMac) Close(fd int) error                              { return nil }
func (u BadMac) Read(fd int, p []byte) (n int, err error)        { return 0, nil }

func TestBadMac(t *testing.T) {
	_, err := Connect("wat", BadMac{})
	if err != ErrorBadMacAddress {
		t.Error("should die if mac address malformed")
	}
}

type NoSocket struct{}

func (u NoSocket) Socket() (int, error)                            { return 0, errors.New("Socket") }
func (u NoSocket) Connect(fd int, addr *unix.SockaddrRFCOMM) error { return errors.New("Connect") }
func (u NoSocket) Close(fd int) error                              { return nil }
func (u NoSocket) Read(fd int, p []byte) (n int, err error)        { return 0, nil }

func TestNoSocket(t *testing.T) {
	_, err := Connect("00:19:10:08:FE:08", NoSocket{})
	if err != ErrorNoSocket {
		t.Error(err)
	}
}

type NoConnection struct{}

func (u NoConnection) Socket() (int, error)                            { return 1, nil }
func (u NoConnection) Connect(fd int, addr *unix.SockaddrRFCOMM) error { return errors.New("Connect") }
func (u NoConnection) Close(fd int) error                              { return nil }
func (u NoConnection) Read(fd int, p []byte) (n int, err error)        { return 0, nil }

func TestNoConnection(t *testing.T) {
	_, err := Connect("00:19:10:08:FE:08", NoConnection{})
	if err != ErrorNoConnection {
		t.Error(err)
	}
}

type Positive struct{}

func (u Positive) Socket() (int, error)                            { return 1, nil }
func (u Positive) Connect(fd int, addr *unix.SockaddrRFCOMM) error { return nil }
func (u Positive) Close(fd int) error                              { return nil }
func (u Positive) Read(fd int, p []byte) (n int, err error)        { return 0, nil }

func TestPositive(t *testing.T) {
	_, err := Connect("00:19:10:08:FE:08", Positive{})
	if err != nil {
		t.Error(err)
	}
}
