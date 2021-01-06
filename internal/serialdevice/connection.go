package serialdevice

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/tzota/nostromo-parker/internal/utils/convert"
	"golang.org/x/sys/unix"
)

var (
	// ErrorBadMacAddress returns when you give bad macaddress
	ErrorBadMacAddress = errors.New("Bad macaddress")

	// ErrorNoSocket returns when unix cant create socket
	ErrorNoSocket = errors.New("Can't create socket")

	// ErrorNoConnection returns when unix cant connect socket to address
	ErrorNoConnection = errors.New("Can't connect socket to address")
)

// Connection represents unix socket data connection
type Connection struct {
	closed     bool
	fd         int
	macAddress string
	unixer     unixSocket
	closer     sdCloser
}

type sdCloser interface {
	Close(fd int) error
}

// unixSocket decouples for testing
type unixSocket interface {
	Socket() (int, error)
	Connect(fd int, addr *unix.SockaddrRFCOMM) error
	Read(fd int, p []byte) (n int, err error)
	// Find my way to decouple things
	sdCloser
	// Close(fd int) error
}

// Connect establishes a connection to the mac address
func Connect(macAddress string, u unixSocket) (Connection, error) {
	conn := Connection{macAddress: macAddress, unixer: u, closer: u}

	// log.SetReportCaller(true)
	mac, err := convert.Str2ba(macAddress)
	if err != nil {
		log.WithField("error", err).Error("Convert mac address")
		return conn, ErrorBadMacAddress
	}

	fd, err := u.Socket()
	if err != nil {
		log.WithField("error", err).Error("Get socket")
		return conn, ErrorNoSocket
	}

	addr := &unix.SockaddrRFCOMM{Addr: mac, Channel: 1}

	log.WithField("mac", macAddress).Info("Connecting")
	err = u.Connect(fd, addr)
	if err != nil {
		log.Error(err)
		return conn, ErrorNoConnection
	}
	log.WithField("mac", macAddress).Info("Connected")
	conn.fd = fd

	return conn, nil
}

// Close closes connection to mac address
func (conn *Connection) Close() error {
	log.WithField("mac", conn.macAddress).Info("Closing connection")
	if conn.closed {
		log.WithField("mac", conn.macAddress).Info("Connection already closed")
		return nil
	}

	err := conn.closer.Close(conn.fd)
	if err == nil {
		conn.closed = true
		log.WithField("mac", conn.macAddress).Info("Closed connection")
	} else {
		log.WithField("mac", conn.macAddress).Info("Can't close connection")
	}

	return err
}

// Read reads bytes from socket
func (conn *Connection) Read(p []byte) (n int, err error) {
	return conn.unixer.Read(conn.fd, p)
}
