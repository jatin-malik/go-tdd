package multicurrency_test

import (
	"testing"

	multicurrency "github.com/jatin-malik/go-tdd/multi-currency"
	"github.com/stretchr/testify/assert"
)

func TestTimes(t *testing.T) {
	// $5 * 2 = $10
	five := multicurrency.Dollar(5)
	assert.Equal(t, multicurrency.Dollar(10), five.Times(2))
	// $5 * 3 = $15
	assert.Equal(t, multicurrency.Dollar(15), five.Times(3))
}

func TestSameCurrencyAddition(t *testing.T) {
	five := multicurrency.Dollar(5)
	three := multicurrency.Dollar(3)
	sum := five.Plus(multicurrency.Bank{}, "USD", three)
	assert.Equal(t, multicurrency.Dollar(8), sum)

}

func TestDiffCurrencyAddition(t *testing.T) {
	five := multicurrency.Dollar(5)
	ten := multicurrency.Pound(10)
	bank := multicurrency.Bank{}
	bank.AddRate("AGP", "USD", 2)
	sum := five.Plus(bank, "USD", ten)
	assert.Equal(t, multicurrency.Dollar(10), sum)

}

func TestMultipleCurrencyAddition(t *testing.T) {
	fiveDollars := multicurrency.Dollar(5)
	tenPounds := multicurrency.Pound(10)
	twentyRupees := multicurrency.Rupee(20)

	bank := multicurrency.Bank{}
	bank.AddRate("AGP", "USD", 2)
	bank.AddRate("INR", "USD", 4)
	money := fiveDollars.Plus(bank, "USD", tenPounds, twentyRupees)
	assert.Equal(t, multicurrency.Dollar(15), money)
}
