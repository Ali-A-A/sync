package single

type Single chan struct{}

// NewSingle generates new Single variable
func NewSingle() Single {
	s := make(chan struct{}, 1)
	s <- struct{}{}
	return s
}

// Do gets f as a function, and if another
// function was already called that, returns
// else it calls the function
func (s Single) Do(f func()) {
	_, ok := <- s
	if !ok {
		return
	}
	f()
	close(s)
}
