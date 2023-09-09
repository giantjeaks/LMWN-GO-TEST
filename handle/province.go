package handle

import (
	"sync"
)

type ProvinceContainer struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *ProvinceContainer) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if name == "" {
		name = "N/A"
	}
	c.counters[name]++
}

func ProvinceCounting(cd *CaseData) map[string]int {
	c := ProvinceContainer{
		counters: map[string]int{},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string) {
		c.inc(name)
		wg.Done()
	}

	wg.Add(len(cd.Data))
	for _, v := range cd.Data {
		go doIncrement(v.ProvinceEn)
	}

	wg.Wait()
	return c.counters
}
