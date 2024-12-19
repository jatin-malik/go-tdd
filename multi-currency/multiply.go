package multicurrency

type Dollar struct {
	amount int
}

func (d *Dollar) times(n int) {
	d.amount *= n
}
