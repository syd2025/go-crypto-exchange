package main

import (
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(2).Second().Do(func() {
		NewKline().Do("1m")
	})
	s.Every(2).Second().Do(func() {
		NewKline().Do("1H")
	})
	s.StartBlocking()
}
