// Alertie
// Copyright (c) 2018 Tobias Urdin
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"github.com/tobias-urdin/alertie/config"
	"errors"
	"os"
	"sync"
	"os/signal"
	"syscall"
	"github.com/nsqio/go-nsq"
	"github.com/tobias-urdin/alertie/log"
)

var c *nsq.Consumer

type MessageHandler struct{}

func (h *MessageHandler) HandleMessage(message *nsq.Message) error {
	if len(message.Body) == 0 {
		return errors.New("Body is empty re-queuing message")
	}

	log.Info("NSQ message received:")
	log.Info(string(message.Body))

	return nil
}

var configFile string

func main() {
	flag.StringVar(&configFile, "config", "alertie.ini", "Path to config file")
	flag.Parse()

	log.Init("worker")
	log.Info("Starting alertie-worker")
	config.Init(configFile)

	//RegisterAlerter("Email", NewEmailAlerter)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	nsqconf := nsq.NewConfig()

	c, err := nsq.NewConsumer("alertie", "alerts", nsqconf)
	if err != nil {
		log.Fatal("Could not create consumer: %s", err)
	}

	// TODO: Make config option
	c.ChangeMaxInFlight(200)

	// TODO: Make number of handlers configurable
	c.AddConcurrentHandlers(
		&MessageHandler{},
		20,
	)

	if err := c.ConnectToNSQLookupds(config.Lookups); err != nil {
		log.Fatal("Could not connect to lookup: %s", err)
	}

	wg.Wait()

	log.Info("Initialization done, now running...")

	//alerter := CreateAlerter("Email")
	//alerter.Alert()

	if c == nil {
		log.Fatal("Failed to initialize NSQ consumer for worker")
		return
        }

	// Lets go
	log.Info("hej2")
	shutdown := make(chan os.Signal, 2)
	signal.Notify(shutdown, syscall.SIGINT)
	log.Info("hej3")

	for {
		select {
		case <-c.StopChan:
			return
		case <-shutdown:
			c.Stop()
		}
	}
}
