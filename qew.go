package main

import (
	"log"
	"time"
)

// Q is the main type to create a queue
type Q struct {
	// functions interface{}
	channel     chan func(...interface{})
	concurrency int
	blocker     chan bool
}

func main() {
	var q Q
	q.concurrency = 1
	q.channel = make(chan func(...interface{}), 100)
	q.blocker = make(chan bool, q.concurrency)

	go q.loop()
	q.Run()

}

func (q *Q) loop() {
	for i := 0; i < 10; i++ {
		// q.channel <- fmt.Sprintf("%d", i)
		go q.Add(log.Println)

	}
}

func (q *Q) Add(f func(...interface{})) {
	q.channel <- f
}

func (q *Q) Run() {
	for {
		q.blocker <- true
		select {
		case str := <-q.channel:

			go q.runner(str)
		}
	}
}

func (q *Q) runner(f func(...interface{})) {
	f("something")
	time.Sleep(time.Second)
	<-q.blocker
}
