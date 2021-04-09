package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type (
	Subscriber chan interface{}
	topicFunc  func(v interface{}) bool
)

type Publisher struct {
	m       sync.RWMutex
	buffer  int
	timeout time.Duration
	subs    map[Subscriber]topicFunc
}

func NewPublisher(timeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:  buffer,
		timeout: timeout,
		subs:    make(map[Subscriber]topicFunc),
	}
}

func (p *Publisher) Subscribe() Subscriber {
	return p.SubscribeTopic(nil)
}

func (p *Publisher) SubscribeTopic(topic topicFunc) Subscriber {
	ch := make(Subscriber, p.buffer)
	p.m.Lock()
	p.subs[ch] = topic
	p.m.Unlock()
	return ch
}

func (p *Publisher) Evict(sub Subscriber) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subs, sub)
	close(sub)
}

func (p *Publisher) sendTopic(message interface{}, sub Subscriber, topic topicFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(message) {
		return
	}

	select {
	case sub <- message:
	case <-time.After(p.timeout):
	}
}

func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	for sub, _ := range p.subs {
		delete(p.subs, sub)
		close(sub)
	}
}

func (p *Publisher) Publish(message interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()
	var wg sync.WaitGroup
	for sub, topic := range p.subs {
		wg.Add(1)
		go p.sendTopic(message, sub, topic, &wg)
	}
	wg.Wait()
}

func main() {
	p := NewPublisher(100*time.Millisecond, 100)

	link1 := p.Subscribe()
	link2 := p.SubscribeTopic(func(message interface{}) bool {
		if s, ok := message.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("gogo")
	p.Publish("Oh my new dream")
	p.Publish("golang is a good language")

	go func() {
		for msg := range link1 {
			fmt.Printf("Link1 %s\n", msg)
		}
	}()

	go func() {
		for msg := range link2 {
			fmt.Printf("Link2 %s\n", msg)
		}
	}()

	time.Sleep(3 * time.Second)

	print("Hello,world")
}
