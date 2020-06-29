package mux

import (
	"encoding/json"
	"errors"
	pool "github.com/h3isenbug/url-shortener/pkg/worker-pool"
)

type Message struct {
	Type string `json:"type"`
}

var (
	ErrHandlerNotFound = errors.New("no handler was registered for this message type")
)

type MessageMux interface {
	SetHandler(messageType string, messageHandler func([]byte) error)
	Handle(messageBytes []byte, after func()) error
}

type MessageMuxV1 struct {
	workerPool pool.WorkerPool
	handlers   map[string]func([]byte) error
}

func NewMessageMuxV1(workerPool pool.WorkerPool) MessageMux {
	return &MessageMuxV1{
		workerPool: workerPool,
		handlers:   make(map[string]func([]byte) error),
	}
}

func (mux MessageMuxV1) SetHandler(messageType string, messageHandler func([]byte) error) {
	mux.handlers[messageType] = messageHandler
}

func (mux MessageMuxV1) Handle(messageBytes []byte, after func()) error {
	var message Message
	if err := json.Unmarshal(messageBytes, &message); err != nil {
		return err
	}

	var handler, found = mux.handlers[message.Type]
	if !found {
		return ErrHandlerNotFound
	}

	return mux.workerPool.AddJob(func() error {
		return handler(messageBytes)
	}, after)
}
