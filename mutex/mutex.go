package mutex

type Mutex struct {
	ch chan struct{}
}

// NewMutex generates new semaphore.
// If n equals to 2, then it creates new binary semaphore
func (m *Mutex) NewMutex(n int) {
	m.ch = make(chan struct{}, n)
}

// Lock acquire the lock
func (m *Mutex) Lock() {
	m.ch <- struct{}{}
}

// Unlock releases the lock
func (m *Mutex) Unlock() {
	<- m.ch
}

// TryLock is similar to Lock but it not block the goroutine
// that called it, and it may cause live lock.
// It returns false if the current goroutine cannot acquire lock.
func (m *Mutex) TryLock() bool {
	select {
	case m.ch <- struct{}{} :
		return true
	default:
		return false
	}
}
