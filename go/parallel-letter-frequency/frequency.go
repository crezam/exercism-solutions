package letter

import "sync"

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(l []string) FreqMap {
	m := FreqMap{}
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	for _, s := range l {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			for _, r := range s {
				mutex.Lock()
				m[r]++
				mutex.Unlock()
			}
		}(s)
	}
	wg.Wait()
	return m
}
