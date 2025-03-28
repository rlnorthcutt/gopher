package core

import (
	"time"

	"github.com/go-co-op/gocron"
)

type JobFunc func(*CoreServices)

type JobScheduler struct {
	scheduler *gocron.Scheduler
	services  *CoreServices
}

func NewJobScheduler(services *CoreServices) *JobScheduler {
	s := gocron.NewScheduler(time.UTC)
	return &JobScheduler{
		scheduler: s,
		services:  services,
	}
}

func (js *JobScheduler) Register(name string, cronExpr string, handler JobFunc) {
	wrapped := func() {
		start := time.Now()
		js.services.Logger.Info().
			Str("job", name).
			Msg("job started")

		defer func() {
			if r := recover(); r != nil {
				js.services.Logger.Error().
					Str("job", name).
					Interface("error", r).
					Msg("job panicked")
			}
		}()

		handler(js.services)

		js.services.Logger.Info().
			Str("job", name).
			Dur("duration", time.Since(start)).
			Msg("job completed")
	}

	_, err := js.scheduler.Cron(cronExpr).Tag(name).Do(wrapped)
	if err != nil {
		js.services.Logger.Error().
			Str("job", name).
			Err(err).
			Msg("failed to register job")
	}
}

func (js *JobScheduler) Start() {
	js.scheduler.StartAsync()
}

func (js *JobScheduler) RunNow(name string) {
	jobs, _ := js.scheduler.FindJobsByTag(name)
	for _, job := range jobs {
		go job.Run()
	}
}
