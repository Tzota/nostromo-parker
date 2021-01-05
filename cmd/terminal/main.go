package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/tzota/nostromo-parker/internal/config"
	"github.com/tzota/nostromo-parker/internal/harvester"
	"github.com/tzota/nostromo-parker/pkg/harvest"
)

func main() {
	cfg, err := config.ReadFromFile("./config.json")
	if err != nil {
		panic(err)
	}

	log.Info("Press Ctrl-C to stop")

	harvest.Simple(cfg, func(message harvester.IMessage) {
		message.ReportToLog()
	})

	select {}
}
