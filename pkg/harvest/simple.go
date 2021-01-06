package harvest

import (
	"os"
	"os/signal"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/tzota/nostromo-parker/internal/config"
	"github.com/tzota/nostromo-parker/internal/dataprovider"
	"github.com/tzota/nostromo-parker/internal/harvester"
	"github.com/tzota/nostromo-parker/internal/serialdevice"
)

// IMessageAction is a simple callback
type IMessageAction = func(harvester.IMessage)

// Simple connects to all the sensors with soe callback
func Simple(cfg config.Config, lambda IMessageAction) {
	var wg sync.WaitGroup

	initOnShutdown(len(cfg.Points))

	for _, point := range cfg.Points {
		if point.Skip {
			continue
		}
		wg.Add(1)
		go func(p config.Point) {
			if err := subscribe(p, lambda); err != nil {
				log.WithFields(log.Fields{"mac": p.Mac, "kind": p.Kind}).Errorf("Can't connect to point")
			}
			wg.Done()
		}(point)
	}
	wg.Wait()
}

func subscribe(p config.Point, lambda IMessageAction) error {
	conn, err := serialdevice.Connect(p.Mac, &serialdevice.UnixSocket{})
	if err != nil {
		return err
	}
	enqueueOnClose(func() { conn.Close() })

	dp := dataprovider.GetChunker(&conn)

	h := p.Kind.New()

	go h.ListenTo(dp)

	go func() {
		for {
			message := <-h.GetMessagesChannel()
			lambda(message)
		}
	}()

	return nil
}

// #region extract to separate structure?

var onCloses []func()

func initOnShutdown(length int) {
	onCloses = make([]func(), 0, length)
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go func(signals <-chan os.Signal) {
		<-signals
		var wg sync.WaitGroup
		for i, onClose := range onCloses {
			wg.Add(1)
			log.Info(i)

			go func(delegate func()) {
				delegate()
				wg.Done()
			}(onClose)
		}
		wg.Wait()
		os.Exit(0)
	}(ch)
}

func enqueueOnClose(f func()) {
	onCloses = append(onCloses, f)
}

// #endregion
