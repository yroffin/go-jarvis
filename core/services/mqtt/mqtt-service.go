// Package mqtt for common interfaces
// MIT License
//
// Copyright (c) 2017 yroffin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package mqtt

import (
	"encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/yroffin/go-boot-sqllite/core/winter"
)

func init() {
	winter.Helper.Register("mqtt-service", (&service{}).New())
}

// service internal members
type service struct {
	// members
	*winter.Service
	// mqtt
	client mqtt.Client
}

// IMqttService implements IBean
type IMqttService interface {
	winter.IService
	PublishExactlyOnce(topic string, message string) error
	PublishLeastOne(topic string, message string) error
	PublishMostOne(topic string, message string) error
	Subscribe(topic string, data interface{}, handler func(interface{}, interface{})) error
}

// New constructor
func (p *service) New() IMqttService {
	bean := service{Service: &winter.Service{Bean: &winter.Bean{}}}
	return &bean
}

// Init this API
func (p *service) Init() error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	opts.SetClientID("jarvis")
	// Listner
	choke := make(chan [2]string)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		choke <- [2]string{msg.Topic(), string(msg.Payload())}
	})
	// Init client
	p.client = mqtt.NewClient(opts)
	if token := p.client.Connect(); token.Wait() && token.Error() != nil {
		log.WithFields(log.Fields{
			"error": token.Error(),
		}).Error("Mqtt - connect")
	}
	// Subcribe
	if token := p.client.Subscribe("#", byte(2), nil); token.Wait() && token.Error() != nil {
		log.WithFields(log.Fields{
			"error": token.Error(),
		}).Error("Mqtt - subscribe")
	}
	// Listner
	go p.listener(choke)
	return nil
}

// PostConstruct this API
func (p *service) PostConstruct(name string) error {
	return nil
}

// Validate this API
func (p *service) Validate(name string) error {
	return nil
}

// Subscribe this API
func (p *service) Subscribe(topic string, data interface{}, handler func(interface{}, interface{})) error {
	// Declare handler
	anonymous := func(client mqtt.Client, msg mqtt.Message) {
		var value interface{}
		json.Unmarshal(msg.Payload(), &value)
		handler(data, value)
	}
	// Subcribe
	if token := p.client.Subscribe(topic, byte(2), anonymous); token.Wait() && token.Error() != nil {
		log.WithFields(log.Fields{
			"error": token.Error(),
		}).Error("Mqtt - subscribe")
	}
	return nil
}

// Mqtt listener
func (p *service) listener(choke chan [2]string) error {
	for {
		incoming := <-choke
		log.WithFields(log.Fields{
			"topic":   incoming[0],
			"message": incoming[1],
		}).Info("Mqtt listener")
	}
}

// Mqtt publisher
func (p *service) publish(topic string, qos int, message string) error {
	p.client.Publish(topic, byte(qos), false, message)
	return nil
}

/**
 * publish a message
 *
 * QoS 2 The highest QoS is 2, it guarantees that each message is received
 * only once by the counterpart. It is the safest and also the slowest
 * quality of service level. The guarantee is provided by two flows there
 * and back between sender and receiver.
 *
 * @param topicName
 * @param payload
 */
func (p *service) PublishExactlyOnce(topic string, message string) error {
	return p.publish(topic, 2, message)
}

/**
 * publish a message
 *
 * QoS 1 – at least once When using QoS level 1, it is guaranteed that a
 * message will be delivered at least once to the receiver. But the message
 * can also be delivered more than once.
 *
 * @param topicName
 * @param payload
 */
func (p *service) PublishLeastOne(topic string, message string) error {
	return p.publish(topic, 1, message)
}

/**
 * publish a message
 *
 * QoS 0 – at most once The minimal level is zero and it guarantees a best
 * effort delivery. A message won’t be acknowledged by the receiver or
 * stored and redelivered by the sender. This is often called “fire and
 * forget” and provides the same guarantee as the underlying TCP protocol.
 *
 * @param topicName
 * @param payload
 */
func (p *service) PublishMostOne(topic string, message string) error {
	return p.publish(topic, 0, message)
}
