package mutex

import (
	"primitives/waitgroup"
	"testing"
)

func TestSemaphore(t *testing.T) {
	var mt Mutex
	mt.NewMutex(3)
	gw := waitgroup.NewGroupWait()
	gw.Add(10)
	sum := 0
	for i := 0; i < 10; i++ {
		go func() {
			defer gw.Done()
			x := mt.TryLock()
			if x {
				sum += 1
			}
		}()
	}
	gw.Wait()
	if sum != 3 {
		t.Fail()
	}
}

func TestTryLock(t *testing.T) {
	var tm Mutex
	tm.NewMutex(1)
	tm.Lock()
	gw := waitgroup.NewGroupWait()
	gw.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer gw.Done()
			x := tm.TryLock()
			if x {
				t.Fail()
			}
		}()
	}
	gw.Wait()
	tm.Unlock()
	x := tm.TryLock()
	if !x {
		t.Fail()
	}
}
