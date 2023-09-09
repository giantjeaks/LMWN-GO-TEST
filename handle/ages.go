package handle

import (
	"covid_sumary/model"
	"sync"
)

type AgeContainer struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *AgeContainer) inc(i *model.Case) {
	c.mu.Lock()
	var name string
	switch {
	case i.Age == nil:
		name = "N/A"
	case *i.Age <= 30:
		name = "0-30"
	case *i.Age > 30 && *i.Age <= 60:
		name = "31-60"
	case *i.Age > 60:
		name = "61+"
	}
	defer c.mu.Unlock()
	c.counters[name]++
}

func AgeCounting(cd *CaseData) map[string]int {
	c := AgeContainer{
		counters: map[string]int{},
	}

	var wg sync.WaitGroup

	doIncrement := func(i *model.Case) {
		c.inc(i)
		wg.Done()
	}

	wg.Add(len(cd.Data))
	for _, v := range cd.Data {

		go doIncrement(&v)

	}

	wg.Wait()
	return c.counters
}
