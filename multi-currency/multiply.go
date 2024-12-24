package multicurrency

type Money struct {
	amount   int
	currency string
}

func (m Money) Times(n int) Money {
	result := m.amount * n
	return Money{result, m.currency}
}

func (augend Money) Plus(bank Bank, to string, addends ...Money) Money {
	amountSum := augend.amount
	for _, addend := range addends {
		amountSum += addend.convertToCurrency(to, bank).amount
	}
	return Money{amountSum, to}
}

func (m Money) convertToCurrency(to string, bank Bank) Money {
	rate := bank.Rate(CurrencyPair{m.currency, to})
	amount := m.amount / rate
	return Money{amount, to}
}
