// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package broker

import (
	"context"
	"sync"
	"time"
)

// Broker broadcasts msgs to registered clients.
type Broker[T any] struct {
	// name of the message broker.
	name string
	// clients is a map of registered clients.
	clients map[chan T]struct{}
	// msgs is the channel for publishing new messages.
	msgs chan T
	// timeout is the timeout for sending a msg to a client.
	timeout time.Duration
	// mutex to synchronize access to the clients map.
	mu sync.Mutex
}

// Default values (assumed, as they were missing in the original code)
const (
	defaultBufferSize = 100
	defaultTimeout    = 5 * time.Second
)

// New creates a new broker.
func New[T any](name string) *Broker[T] {
	return &Broker[T]{
		clients: make(map[chan T]struct{}),
		msgs:    make(chan T, defaultBufferSize),
		timeout: defaultTimeout,
		name:    name,
	}
}

// Name returns the name of the msg broker.
func (b *Broker[T]) Name() string {
	return b.name
}

// Start starts the broker loop.
func (b *Broker[T]) Start(ctx context.Context) error {
	go b.start(ctx)
	return nil
}

// start starts the broker loop.
func (b *Broker[T]) start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// close all leftover clients and break the broker loop
			b.mu.Lock()
			for client := range b.clients {
				b.Unsubscribe(client)
			}
			b.mu.Unlock()
			return
		case msg := <-b.msgs:
			// broadcast published msg to all clients
			b.mu.Lock()
			for client := range b.clients {
				// send msg to client (or discard msg after timeout)
				select {
				case client <- msg:
				case <-time.After(b.timeout):
				}
			}
			b.mu.Unlock()
		}
	}
}

// Publish publishes a msg to the broker.
// Returns context error on timeout or cancel.
func (b *Broker[T]) Publish(ctx context.Context, msg T) error {
	select {
	case b.msgs <- msg:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Subscribe registers a new client to the broker and returns it to the caller.
func (b *Broker[T]) Subscribe() (chan T, error) {
	client := make(chan T)
	b.mu.Lock()
	b.clients[client] = struct{}{}
	b.mu.Unlock()
	return client, nil
}

// Unsubscribe removes a client from the broker.
func (b *Broker[T]) Unsubscribe(client chan T) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if _, ok := b.clients[client]; ok {
		delete(b.clients, client)
		close(client)
	}
}

