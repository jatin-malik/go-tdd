package multicurrency

type CurrencyPair struct {
	cur1 string
	cur2 string
}

type Bank struct {
	exchangeRates map[CurrencyPair]int
}

func (b Bank) Rate(curPair CurrencyPair) int {
	if curPair.cur1 == curPair.cur2 {
		return 1
	}
	return b.exchangeRates[curPair]
}

func (b *Bank) AddRate(from string, to string, rate int) {
	if b.exchangeRates == nil {
		b.exchangeRates = make(map[CurrencyPair]int)
	}

	b.exchangeRates[CurrencyPair{from, to}] = rate
}
