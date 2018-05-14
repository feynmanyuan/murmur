/*
 * Copyright (c) 2018. feynmanyuan@gmail.com All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0;
 *
 */
package pubsub

import (
	"sync"
)

type MessageAcceptor func(interface{}) bool

type Publisher struct {
	channels 	[]*channel
	lock   		*sync.RWMutex
}

type channel struct {
	pred 	MessageAcceptor
	ch 		chan interface{}
}

func NewPublisher() *Publisher {
	return &Publisher{
		channels:		make([]*channel, 0),
		lock:			&sync.RWMutex{},
	}
}

func (pub *Publisher)Sub(acceptor MessageAcceptor) chan interface{} {
	pub.lock.Lock()
	defer pub.lock.Unlock()

	c := &channel{ch:make(chan interface{}, 10), pred:acceptor}

	pub.channels = append(pub.channels, c)

	return c.ch
}

func (pub *Publisher)Publish(msg interface{}) error {
	defer func() {
		recover()
	}()

	pub.lock.RLock()
	channels := pub.channels
	pub.lock.RUnlock()

	for _, c := range channels {
		if c.pred(msg) {
			c.ch <- msg
		}
	}
	return nil
}