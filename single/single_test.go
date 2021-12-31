package single

import (
	"sync"
	"testing"
)

func TestOnce(t *testing.T) {
	s :=NewSingle()
	sum := 0
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			s.Do(func() {
				sum += 1
			})
		}()
	}
	wg.Wait()
	if sum != 1 {
		t.Errorf("sum: %v", sum)
	}
}
