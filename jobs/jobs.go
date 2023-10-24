package jobs

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/just-arun/micro-auth/model"
)

type cronJob struct {
	ctx  *model.HandlerCtx
	cron *gocron.Scheduler
}

func Register(ctx *model.HandlerCtx) {
	s := gocron.NewScheduler(time.UTC)

	c := &cronJob{ctx: ctx, cron: s}

	c.DeleteRole()

	s.StartAsync()
}
