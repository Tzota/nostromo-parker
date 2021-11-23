package harvest

import (
	"os"
	"os/signal"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tzota/nostromo-parker/internal/config"
	"github.com/tzota/nostromo-parker/internal/dataprovider"
	"github.com/tzota/nostromo-parker/internal/harvester"
	"github.com/tzota/nostromo-parker/internal/serialdevice"
	"github.com/tzota/nostromo-parker/internal/utils/mymath"
)

// IMessageAction is a simple callback
type IMessageAction = func(harvester.IMessage)

// Simple connects to all the sensors with soe callback
func Simple(cfg config.Config, lambda IMessageAction) {
	var wg sync.WaitGroup

	initOnShutdown(len(cfg.Points))

	for _, point := range cfg.Points {
		if point.Skip {
			log.WithField("mac", point.Mac).Info("Skipping point")
			continue
		}

		firstRun := true
		timeOuts := make(map[string]int64)

		wg.Add(1)
		go func(p config.Point) {
			timeOuts[p.Mac] = 1
			go (func() {
				for {
					err, flag := subscribe(p, lambda)
					if firstRun {
						firstRun = false
						wg.Done()
					}
					if err != nil {
						timeout := mymath.Int64min(128, timeOuts[p.Mac]*2)
						timeOuts[p.Mac] = timeout
						log.WithFields(log.Fields{"mac": p.Mac, "kind": p.Kind, "timeout": timeout}).Errorf("Can't connect to point")
						time.Sleep(time.Second * time.Duration(timeout))
					} else {
						timeOuts[p.Mac] = 1
					}
					if flag != nil {
						<-flag
						log.WithFields(log.Fields{"mac": p.Mac, "kind": p.Kind}).Errorf("Lost connection to point")
					}
				}
			})()
		}(point)
	}
	wg.Wait()
}

func subscribe(p config.Point, lambda IMessageAction) (error, chan bool) {
	conn, err := serialdevice.Connect(p.Mac, &serialdevice.UnixSocket{})
	if err != nil {
		return err, nil
	}
	enqueueOnClose(func() { conn.Close() })

	dp, flag := dataprovider.GetChunker(&conn)

	h := p.Kind.New()

	go h.ListenTo(dp)

	go func() {
		for {
			message := <-h.GetMessagesChannel()
			lambda(message)
		}
	}()

	return nil, flag
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
