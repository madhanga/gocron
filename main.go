package main

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

var count int

type Greeter struct{}

func (g *Greeter) Run() {
	log.Println("hello")
}

type Faulty struct{}

func (f *Faulty) Run() {
	panic("failing faulty")
}

func main() {
	c := cron.New(cron.WithChain(
		cron.Recover(cron.DefaultLogger),
	))

	var greeter Greeter
	_, _ = c.AddJob("@every 1s", &greeter)

	var faulty Faulty
	_, _ = c.AddJob("@every 5s", &faulty)

	c.Start()
	defer c.Stop()

	time.Sleep(10 * time.Minute)
}
