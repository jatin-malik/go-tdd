package multicurrency

// All supported currencies by the package should be listed here.

func Dollar(amount int) Money {
	return Money{amount, "USD"}
}

func Pound(amount int) Money {
	return Money{amount, "AGP"}
}

func Rupee(amount int) Money {
	return Money{amount, "INR"}
}
