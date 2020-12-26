package engine

import (
	"sync"
)

type messageQueue struct {
	sync.Mutex
	data             []Command
	receiveSignal    chan struct{}
	receiveRequested bool
}

type EventLoop struct {
	mq          *messageQueue
	stopSignal  chan struct{}
	stopRequest bool
}

func (el *EventLoop) Start() {
	el.mq = &messageQueue{receiveSignal: make(chan struct{})}
	el.stopSignal = make(chan struct{})
	go func() {
		for !el.stopRequest || len(el.mq.data) != 0 {
			cmd := el.mq.pull()
			cmd.Execute(el)
		}
		el.stopSignal <- struct{}{}
	}()
}

func (el *EventLoop) Post(cmd Command) {
	el.mq.push(cmd)
}

type CommandFunc func(h Handler)

func (cf CommandFunc) Execute(h Handler) {
	cf(h)
}

func (el *EventLoop) AwaitFinish() {
	el.Post(CommandFunc(func(h Handler) {
		el.stopRequest = true
	}))
	<-el.stopSignal
}

func (mq *messageQueue) push(command Command) {
	mq.Lock()
	defer mq.Unlock()

	mq.data = append(mq.data, command)
	if mq.receiveRequested {
		mq.receiveRequested = false
		mq.receiveSignal <- struct{}{}
	}

}

func (mq *messageQueue) pull() Command {
	mq.Lock()
	defer mq.Unlock()

	if len(mq.data) == 0 {
		mq.receiveRequested = true
		mq.Unlock()
		<-mq.receiveSignal
		mq.Lock()
	}

	res := mq.data[0]
	mq.data[0] = nil
	mq.data = mq.data[1:]
	return res
}
