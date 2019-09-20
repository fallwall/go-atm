package funding

type Fund struct {
	// balance is unexported (private), because it's lowercase
	balance int
}

// A regular function returning a pointer to a fund
func NewFund(initialBalance int) *Fund {
	// We can return a pointer to a new struuct without worrying about
	// whether it's ont he stack or heap: Go figres that out for s
	return &Fund {
		balance: initialBalance,
	}
}

// Methods start with a *receriver*, in this case a Fund pointer
func (f * Fund) Balance() int {
	return f.balance
}

func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}