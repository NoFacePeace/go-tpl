package scrape

import "context"

func NewScrape(ctx context.Context) []Dispatch {
	scraperScheduler := NewScheduler(ctx, nil)
	return []Dispatch{
		scraperScheduler,
	}
}

type Scraper interface {
	Match(ctx context.Context) bool
	Produce(ctx context.Context)
}
