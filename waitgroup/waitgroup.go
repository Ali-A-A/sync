package waitgroup

type gen struct {
	wait chan struct{}
	n int
}

func newGen() gen {
	return gen{ wait: make(chan struct{}), n: 0}
}

func (g gen) end() {
	close(g.wait)
}

type GroupWait chan gen

// NewGroupWait generates new groupWait variable
func NewGroupWait() GroupWait {
	wg := make(GroupWait, 1)
	g := newGen()
	g.end()
	wg <- g
	return wg
}

// Add increases or decreases the number that
// Wait should waits for them.
func (wg GroupWait) Add(delta int) {
	g := <-wg
	if g.n == 0 {
		g = newGen()
	}
	g.n += delta
	if g.n < 0 {
		panic("negative GroupWait count")
	}
	if g.n == 0 {
		g.end()
	}
	wg <- g
}

// Done just calls Add with -1 input
func (wg GroupWait) Done() { wg.Add(-1) }

// Wait is waits until all goroutines call Done
func (wg GroupWait) Wait() {
	g := <-wg
	wait := g.wait
	wg <- g
	<-wait
}
