package engine

import (
	"sync"
)

type EventLoop struct {
	messageQueue chan Command
	wg sync.WaitGroup
	count int
	done int
}

func (el *EventLoop) readMessages(){
	for i := range el.messageQueue{
		i.Execute(el)
		el.done++
		if el.count == el.done { el.wg.Done() }
	}

}

func (el *EventLoop) Start() {
	el.messageQueue = make(chan Command, 1)
	el.wg.Add(1)
	go el.readMessages()
}

func (el *EventLoop) Post(cmd Command) {
	el.count++
	go func() {
		el.messageQueue <- cmd
	}()
}

func (el *EventLoop) AwaitFinish() {
	defer close(el.messageQueue)
	el.wg.Wait()
}