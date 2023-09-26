package main

import (
	"fmt"
	"sync"
)

type EventEmitterX struct {
	subscribers map[string][]chan interface{}
	mu          sync.Mutex
}


func SingletonEmitter() *EventEmitterX{
	return &EventEmitterX{
		subscribers: make(map[string][]chan interface{}),
	}
}

func(emitter *EventEmitterX) On(event string) chan interface{}{
	emitter.mu.Lock()
	defer emitter.mu.Unlock()

	ch := make(chan interface{})
	emitter.subscribers[event] = append(emitter.subscribers[event], ch)

	return ch
}

func (emitter *EventEmitterX) Emit(event string, data interface{}){
	emitter.mu.Lock()
	defer emitter.mu.Unlock();

	subscribers, exists := emitter.subscribers[event]
	if !exists {
		return
	}
	for _, ch := range subscribers {
		go func(ch chan interface{}) {
			ch <- data
		}(ch)
	}
}

func main(){
	emitter := SingletonEmitter()

	// Listen to events
	subscriber1 := emitter.On("event1");
	subscriber2 := emitter.On("event2");

	// Publish events
	emitter.Emit("event1", "Hello from event 1!")
	emitter.Emit("event2", "Greetings from event 2!")

	// Receive and print events
	fmt.Println("Event 1:", <-subscriber1)
	fmt.Println("Event 2:", <-subscriber2)
}