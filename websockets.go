package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"time"
)

var brokerLoad = make(chan bool)
var brokerConnection = make(chan bool)
var brokerClients = make(chan bool)

func brokerLoadHandler(client MQTT.Client, msg MQTT.Message) {
	brokerLoad <- true
	fmt.Printf("BrokerLoadHandler         ")
	fmt.Printf("[%s]  ", msg.Topic())
	fmt.Printf("%s\n", msg.Payload())
}

func brokerConnectionHandler(client MQTT.Client, msg MQTT.Message) {
	brokerConnection <- true
	fmt.Printf("BrokerConnectionHandler   ")
	fmt.Printf("[%s]  ", msg.Topic())
	fmt.Printf("%s\n", msg.Payload())
}

func brokerClientsHandler(client MQTT.Client, msg MQTT.Message) {
	brokerClients <- true
	fmt.Printf("BrokerClientsHandler")
	fmt.Printf("[%s]  ", msg.Topic())
	fmt.Printf("%s\n", msg.Payload())
}

func main() {
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://iot.eclipse.org:1883")
	opts.SetClientID("custom-store2")

	var callback MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Printf("MSG: %s\n", msg.Payload())
	}

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	c.Subscribe("/go-mqtt/sample", 0, callback)


	time.Sleep(100000 * time.Millisecond)
	c.Disconnect(250)
}