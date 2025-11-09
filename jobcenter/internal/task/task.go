package task

import (
	"jobcenter/internal/kline"
	"jobcenter/internal/svc"
	"time"

	"github.com/go-co-op/gocron"
)

type Task struct {
	s   *gocron.Scheduler
	ctx *svc.ServiceContext
}

func NewTask(ctx *svc.ServiceContext) *Task {
	return &Task{
		s:   gocron.NewScheduler(time.UTC),
		ctx: ctx,
	}
}

func (t *Task) Run() {
	t.s.Every(2).Second().Do(func() {
		kline.NewKline(t.ctx.Config.Okx, t.ctx.MongoClient, t.ctx.KafkaClient).Do("1m")
	})
	t.s.Every(2).Second().Do(func() {
		kline.NewKline(t.ctx.Config.Okx, t.ctx.MongoClient, t.ctx.KafkaClient).Do("1H")
	})
}

func (t *Task) Stop() {
	t.s.Stop()
}

func (t *Task) StartBlocking() {
	t.s.StartBlocking()
}
