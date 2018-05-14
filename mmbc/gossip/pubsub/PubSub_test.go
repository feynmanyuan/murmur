/*
 * Copyright (c) 2018. feynmanyuan@gmail.com All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0;
 *
 */
package pubsub

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"fmt"
	"time"
)

func TestPublisher_Publish(t *testing.T) {
	publisher := NewPublisher()
	ch := publisher.Sub(func(i interface{}) bool {
		return true
	})

	publisher.Publish("Test")

	msg := <-ch

	assert.Equal(t, "Test", msg)
}

func TestPublisher_Sub(t *testing.T) {
	publisher := NewPublisher()

	ch1 := publisher.Sub(func(i interface{}) bool {
		return true
	})
	ch2 := publisher.Sub(func(i interface{}) bool {
		return true
	})

	publisher.Publish(1)

	total := int32(0)

	closeFlag := make(chan bool, 0)

	go func() {
		for {
			select{
			case m1 := <-ch1:
				if m, ok := m1.(int); ok {
					fmt.Printf("test1 - %d\n", m)
					closeFlag <- true
				}
			case m2 := <-ch2:
				if m, ok := m2.(int); ok {
					fmt.Printf("test2 - %d\n", m)
					closeFlag <- true
				}
			}
		}
	}()

	for {
		select{
		case <-closeFlag:
			fmt.Println("Close Flag")
			atomic.AddInt32(&total, int32(1))
		}

		if total >= 2 {
			break
		}
	}
	assert.Equal(t, int32(2), total)
}

func TestPublisher_Sub2(t *testing.T) {
	publisher := NewPublisher()

	ch1 := publisher.Sub(func(i interface{}) bool {
		return true
	})
	ch2 := publisher.Sub(func(i interface{}) bool {
		return false
	})

	publisher.Publish(2)

	total := int32(0)

	closeFlag := make(chan bool, 0)

	go func() {
		for {
			select{
			case m1 := <-ch1:
				if m, ok := m1.(int); ok {
					fmt.Printf("test1 - %d\n", m)
					closeFlag <- true
				}
			case m2 := <-ch2:
				if m, ok := m2.(int); ok {
					fmt.Printf("test2 - %d\n", m)
					closeFlag <- true
				}
			}
		}
	}()

	isClose := false
	for {
		select{
		case <-closeFlag:
			fmt.Println("Close Flag")
			atomic.AddInt32(&total, int32(1))
		case <- time.After(time.Second):
			fmt.Println("time out")
			isClose = true
		}
		if isClose {
			break
		}
	}
	assert.Equal(t, int32(1), total)
}