package main

import (
	"sync"
	"github.com/streadway/amqp"
)

/**
 * connect 结构体
 */
type MqConnection struct {
	Lock       sync.RWMutex
	Connection *amqp.Connection
	MqUri      string
}

type ChannelContext struct {
	Exchange     string
	ExchangeType string
	RoutingKey   string
	Reliable     bool
	Durable      bool
	ChannelId    string
	Channel      *amqp.Channel
}

type BaseMq struct {
	MqConnection *MqConnection

	//channel cache
	ChannelContexts map[string]*ChannelContext
}

func (bmq *BaseMq) Init() {
	bmq.ChannelContexts = make(map[string]*ChannelContext)
}
