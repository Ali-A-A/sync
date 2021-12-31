package waitgroup

import (
	"primitives/mutex"
	"testing"
)

func TestGroupWait(t *testing.T) {
	gw := NewGroupWait()
	var tm mutex.Mutex
	tm.NewMutex(1)
	gw.Add(10)
	sum := 0
	for i := 0; i < 10; i++ {
		go func() {
			defer gw.Done()
			tm.Lock()
			sum += 1
			tm.Unlock()
		}()
	}
	gw.Wait()
	if sum != 10 {
		t.Fail()
	}
}
