package scrape

import (
	"context"
	"time"

	"log"

	"github.com/NoFacePeace/go-tpl/goroutine"
	"github.com/NoFacePeace/go-tpl/locker"
	"github.com/NoFacePeace/go-tpl/times"
)

var maxWaitNoMessage = time.Minute * 15

const (
	keyScheduleRestart = "" // 调度重启分布式锁
	keyMarkSchedule    = "" // 上一次调度时间
	keyLocker          = ""
)

func NewScheduler(ctx context.Context, scrapers []Scraper) *Scheduler {
	s := &Scheduler{
		scrapers:   scrapers,
		nextTicker: time.NewTicker(maxWaitNoMessage),
	}
	s.WatchAndNext(ctx)
	now := time.Now()
	s.Produce(ctx, now, times.ZeroOutInSecond(now).Add(time.Minute), 0)
	return s
}

type Scheduler struct {
	scrapers   []Scraper
	nextTicker *time.Ticker
}

func (s *Scheduler) Router() string {
	return "scheduler"
}

func (s *Scheduler) Consume(ctx context.Context, msg []byte) error {
	Scheduled, err := time.Parse(time.RFC3339, string(msg))
	if err != nil {
		log.Println(err)
	}
	if times.ZeroOutInSecond(time.Now()).Sub(times.ZeroOutInSecond(Scheduled)) < time.Minute {
		s.Produce(ctx, Scheduled, times.ZeroOutInSecond(Scheduled).Add(time.Minute), 0)
		return nil
	}
	ok := locker.Lock(ctx, keyLocker, time.Second*40)
	if !ok {
		return nil
	}
	defer locker.Unlock(ctx, keyLocker)
	if err := s.checkAndResetTicker(ctx, Scheduled); err != nil {
		return nil
	}
	for _, scraper := range s.scrapers {
		if ok := scraper.Match(ctx); ok {
			scraper.Produce(ctx)
		}
	}
	s.Produce(ctx, time.Now(), time.Time{}, 0)
	return nil
}

func (s *Scheduler) WatchAndNext(ctx context.Context) {
	goroutine.SafeGo(ctx, func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-s.nextTicker.C:
				// 获取上一次调度时间
				lastScheduled := s.getScheduled()
				// 对比现在时间，如果小于 2 分钟就 continue
				if !lastScheduled.IsZero() && time.Since(lastScheduled) < 2*time.Minute {
					continue
				}
				// 大于 2 分钟，抢占 keyScheduleRestart 分布式锁
				ok := locker.Lock(ctx, keyScheduleRestart, time.Minute)
				if !ok {
					continue
				}
				// 生产消息
				now := time.Now()
				s.Produce(ctx, now, times.ZeroOutInSecond(now).Add(time.Minute), 0)
			}
		}
	})
}

func (s *Scheduler) Produce(ctx context.Context, scheduled, deliverAt time.Time, deliverAfter time.Duration) {

}

func (s *Scheduler) checkAndResetTicker(ctx context.Context, scheduled time.Time) error {
	// 获取 keyMarkSchedule
	lastScheduled := s.getScheduled()
	if !lastScheduled.IsZero() && (times.BeforeInMin(scheduled, lastScheduled) || times.EqualInMin(scheduled, lastScheduled)) {
		return nil
	}
	// 重置 ticker
	s.nextTicker.Reset(maxWaitNoMessage)
	// 设置 keyMarkSchedule
	s.setScheduled(scheduled)
	return nil
}

func (s *Scheduler) getScheduled() time.Time {
	return time.Now()
}

func (s *Scheduler) setScheduled(t time.Time) {
}
