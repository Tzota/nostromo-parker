package serialdevice

import (
	"errors"
	"testing"
)

func TestCloseClosed(t *testing.T) {
	c := Connection{closed: true}
	err := c.Close()
	if err != nil {
		t.Error(err)
	}
}

type UnClosable struct{}

// func (u UnClosable) Socket() (int, error)                            { return 1, nil }
// func (u UnClosable) Connect(fd int, addr *unix.SockaddrRFCOMM) error { return nil }
func (u UnClosable) Close(fd int) error { return errors.New("") }

// func (u UnClosable) Read(fd int, p []byte) (n int, err error)        { return 1, nil }

func TestCloseUnclosable(t *testing.T) {
	c := Connection{closed: false, closer: UnClosable{}}
	err := c.Close()
	if err == nil {
		t.Error(err)
	}
	if c.closed {
		t.Error("connection should not be closed after exception")
	}
}

type Closable struct{}

// func (u Closable) Socket() (int, error)                            { return 1, nil }
// func (u Closable) Connect(fd int, addr *unix.SockaddrRFCOMM) error { return nil }
func (u Closable) Close(fd int) error { return nil }

// func (u Closable) Read(fd int, p []byte) (n int, err error)        { return 1, nil }

func TestClosePositive(t *testing.T) {
	c := Connection{closed: false, closer: Closable{}}
	err := c.Close()
	if err != nil {
		t.Error(err)
	}
	if !c.closed {
		t.Error("connection was not closed")
	}
}
