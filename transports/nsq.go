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

package transports

import (
	"fmt"
	"errors"
	"gopkg.in/ini.v1"
	"github.com/nsqio/go-nsq"
)

const DEFAULT_CONCURRENCY = 10

type NsqTransport struct {
	lookups string
	topic string
	channel string
	concurrency int
	config *nsq.Config
	producer *nsq.Producer
	consumer *nsq.Consumer
}

type messageHandler struct {}

func (m messageHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Printf("Body: %s", msg.Body)
	msg.Finish()
	return nil
}

func (t NsqTransport) ParseConfig(config *ini.Section) error {
	if !config.HasKey("lookups") {
		return errors.New("NSQ no lookups in config")
	}
	key, _ := config.GetKey("lookups")
	t.lookups = key.String()

	if !config.HasKey("topic") {
		return errors.New("NSQ no topic in config")
	}
	key, _ = config.GetKey("topic")
	t.topic = key.String()

	if !config.HasKey("channel") {
		return errors.New("NSQ no channel in config")
	}
	key, _ = config.GetKey("channel")
	t.channel = key.String()

	if config.HasKey("concurrency") {
		c, _ := config.GetKey("concurrency")
		t.concurrency = c.MustInt(DEFAULT_CONCURRENCY)
	} else {
		t.concurrency = DEFAULT_CONCURRENCY
	}

	return nil
}

func (t NsqTransport) InitProducer(config *ini.Section) error {
	err := t.ParseConfig(config)
	if err != nil {
		return err
	}
	t.config = nsq.NewConfig()
	t.producer, err = nsq.NewProducer(t.lookups, t.config)
	return err
}

func (t NsqTransport) InitConsumer(config *ini.Section) error {
	err := t.ParseConfig(config)
	if err != nil {
		return err
	}
	t.config = nsq.NewConfig()
	t.consumer, err = nsq.NewConsumer(t.topic, t.channel, t.config)
	if err != nil {
		return err
	}
	return err
}

func (t NsqTransport) Publish(data []byte) error {
	if t.producer == nil {
		panic("NSQ transport producer is nil")
	}
	return t.producer.Publish(t.topic, data)
}

func (t NsqTransport) Consume() error {
	if t.consumer == nil {
		panic("NSQ transport consumer is nil")
	}
	t.consumer.AddConcurrentHandlers(&messageHandler{}, t.concurrency)
	err := t.consumer.ConnectToNSQLookupd(t.lookups)
	return err
}
