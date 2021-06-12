package app

import (
	"log"

	"github.com/jasonlvhit/gocron"
)

type Scheduler struct {
	Service *Service
}

func (s *Scheduler) Schedule(number uint64, timeframe string, task interface{}) {
	log.Printf("Schedule agent in every %v %v\n", number, timeframe)

	switch timeframe {
	case Second:
		gocron.Every(number).Seconds().Do(task)
	case Minute:
		gocron.Every(number).Minutes().Do(task)
	case Hour:
		gocron.Every(number).Hours().Do(task)
	case Day:
		gocron.Every(number).Days().Do(task)
	case Week:
		gocron.Every(number).Weeks().Do(task)
	}
	_, time := gocron.NextRun()
	log.Printf("- execute at: %v\n", time)
}

func (s *Scheduler) ScheduleEvery(timeframe string, time string, task interface{}) {
	log.Printf("Schedule agent in every %v on %v\n", timeframe, time)

	switch timeframe {
	case Day:
		gocron.Every(1).Day().At(time).Do(task)
	case Monday:
		gocron.Every(1).Monday().At(time).Do(task)
	case Tuesday:
		gocron.Every(1).Tuesday().At(time).Do(task)
	case Wednesday:
		gocron.Every(1).Wednesday().At(time).Do(task)
	case Thursday:
		gocron.Every(1).Thursday().At(time).Do(task)
	case Friday:
		gocron.Every(1).Friday().At(time).Do(task)
	case Saturday:
		gocron.Every(1).Saturday().At(time).Do(task)
	case Sunday:
		gocron.Every(1).Sunday().At(time).Do(task)
	}
}

func (s *Scheduler) Start() {
	<-gocron.Start()
}
