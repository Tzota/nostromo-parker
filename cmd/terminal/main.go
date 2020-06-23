package main

import (
	"syscall"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"

	"github.com/tzota/nostromo-parker/internal/dataprovider"
	"github.com/tzota/nostromo-parker/internal/sensors/ds18b20"
	"github.com/tzota/nostromo-parker/internal/utils/convert"
)

func main() {
	// handle := serialdevice.Connect(mac)
	// defer := handle.Close();
	// dp := dataprovider.GetChunker(dataprovider.RealUnixReader{}, handle)
	// harvester := ds18b20.New()
	// go harvester.ListenTo(dp)

	mac, err := convert.Str2ba("00:19:10:08:FE:08")
	if err != nil {
		log.Fatal(err)
	}

	fd, err := unix.Socket(syscall.AF_BLUETOOTH, syscall.SOCK_STREAM, unix.BTPROTO_RFCOMM)
	check(err)
	addr := &unix.SockaddrRFCOMM{Addr: mac, Channel: 1}

	log.Print("connecting...")
	err = unix.Connect(fd, addr)
	check(err)
	defer unix.Close(fd)
	log.Println("done")

	dp := dataprovider.GetChunker(dataprovider.RealUnixReader{}, fd)
	harvester := ds18b20.New()

	go harvester.ListenTo(dp)

	go func() {
		for {
			message := <-harvester.Messages

			sendToTerminal(message)
		}
	}()

	select {}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func sendToTerminal(m ds18b20.Message) {
	log.WithFields(log.Fields{
		"message": m.Temperature,
	}).Info("Received")
}
