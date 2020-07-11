package main

import (
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"

	"github.com/tzota/nostromo-parker/internal/config"
	"github.com/tzota/nostromo-parker/internal/dataprovider"
	"github.com/tzota/nostromo-parker/internal/serialdevice"
)

func init() {
	// log.SetReportCaller(true)
}

func main() {
	cfg, err := config.ReadFromFile("./config.json")
	if err != nil {
		panic(err)
	}

	for _, point := range cfg.Points {
		if point.Skip {
			continue
		}
		if err = subscribe(point); err != nil {
			log.WithFields(log.Fields{
				"mac": point.Mac, "kind": point.Kind,
			}).Errorf("Can't connect to point")
		}
	}

	select {}
}

func subscribe(p config.Point) error {
	conn, err := serialdevice.Connect(p.Mac, &serialdevice.UnixSocket{})
	if err != nil {
		return err
	}
	subscribeShutdown(func() { conn.Close() })

	dp := dataprovider.GetChunker(&conn)

	h := p.Kind.New()

	go h.ListenTo(dp)

	go func() {
		for {
			message := <-h.GetMessagesChannel()
			message.ReportToLog()
		}
	}()

	return nil
}

func subscribeShutdown(onClose func()) {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go func(signals <-chan os.Signal) {
		<-signals
		onClose()
		os.Exit(0)
	}(ch)
}
