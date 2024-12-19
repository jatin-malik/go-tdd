package multicurrency

import "testing"

func TestMultiply(t *testing.T) {
	// $5 * 2 = $10
	fiveDollars := Dollar{5}
	fiveDollars.times(2)
	if fiveDollars.amount != 10 {
		t.Errorf("$5*2=%d, wanted $10", fiveDollars.amount)
	}
}
