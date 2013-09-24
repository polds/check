package check

import (
	"testing"
)

func TestAny(t *testing.T) {
	var small, fifteen, sixteen string

	// Check for a small number
	small = "4111"
	if CreditCard(small).CardIs("Any") {
		t.Errorf("Any(%v) = true, want false", small)
	}

	// Check for a AMEX style card number
	fifteen = "822233334444555"
	if !CreditCard(fifteen).CardIs("Any") {
		t.Errorf("Any(%v) = false, want true", fifteen)
	}

	// Check for all other style card numbers
	sixteen = "9111222233334444"
	if !CreditCard(sixteen).CardIs("Any") {
		t.Errorf("Any(%v) = false, want true", sixteen)
	}

}

func TestMasterCard(t *testing.T) {
	var small, fifteen, sixteen string

	// Check for a small number
	small = "4111"
	if CreditCard(small).CardIs("MasterCard") {
		t.Errorf("Any(%v) = true, want false", small)
	}

	// Check for a AMEX style card number (numbers)
	fifteen = "822233334444555"
	if CreditCard(fifteen).CardIs("MasterCard") {
		t.Errorf("Any(%v) = true, want false", fifteen)
	}

	// Check for MasterCard card number
	sixteen = "5011122223333444"
	if !CreditCard(sixteen).CardIs("MasterCard") {
		t.Log(len(sixteen))
		t.Errorf("Any(%v) = false, want true", sixteen)
	}

	// Check for all other style card numbers
	sixteen = "9111222233334444"
	if CreditCard(sixteen).CardIs("MasterCard") {
		t.Errorf("Any(%v) = true, want false", sixteen)
	}
}

func TestVisa(t *testing.T) {
	var small, fifteen, sixteen string

	// Check for a small number
	small = "4111"
	if CreditCard(small).CardIs("Visa") {
		t.Errorf("Any(%v) = true, want false", small)
	}

	// Check for a AMEX style card number (numbers)
	fifteen = "822233334444555"
	if CreditCard(fifteen).CardIs("Visa") {
		t.Errorf("Any(%v) = true, want false", fifteen)
	}

	// Check for Visa card number
	sixteen = "4011122223333444"
	if !CreditCard(sixteen).CardIs("Visa") {
		t.Log(len(sixteen))
		t.Errorf("Any(%v) = false, want true", sixteen)
	}

	// Check for all other style card numbers
	sixteen = "9111222233334444"
	if CreditCard(sixteen).CardIs("Visa") {
		t.Errorf("Any(%v) = true, want false", sixteen)
	}
}

func TestDiscover(t *testing.T) {
	var small, fifteen, sixteen string

	// Check for a small number
	small = "4111"
	if CreditCard(small).CardIs("Discover") {
		t.Errorf("Any(%v) = true, want false", small)
	}

	// Check for a AMEX style card number (numbers)
	fifteen = "822233334444555"
	if CreditCard(fifteen).CardIs("Discover") {
		t.Errorf("Any(%v) = true, want false", fifteen)
	}

	// Check for Discover card number
	sixteen = "6011122223333444"
	if !CreditCard(sixteen).CardIs("Discover") {
		t.Log(len(sixteen))
		t.Errorf("Any(%v) = false, want true", sixteen)
	}

	// Check for all other style card numbers
	sixteen = "9111222233334444"
	if CreditCard(sixteen).CardIs("Discover") {
		t.Errorf("Any(%v) = true, want false", sixteen)
	}
}

func TestAmex(t *testing.T) {
	var small, fifteen, sixteen string

	// Check for a small number
	small = "4111"
	if CreditCard(small).CardIs("AmericanExpress") {
		t.Errorf("Any(%v) = true, want false", small)
	}

	// Check for a AMEX style card number (numbers)
	fifteen = "822233334444555"
	if CreditCard(fifteen).CardIs("AmericanExpress") {
		t.Errorf("Any(%v) = true, want false", fifteen)
	}

	// Check for AMEX style card number (sequence) Format 1
	fifteen = "3417122223333444"
	if !CreditCard(fifteen).CardIs("AmericanExpress") {
		t.Log(len(fifteen))
		t.Log("Failure format 1")
		t.Errorf("Any(%v) = false, want true", fifteen)
	}

	// Check for AMEX style card number (sequence) Format 2
	fifteen = "3717122223333444"
	if !CreditCard(fifteen).CardIs("AmericanExpress") {
		t.Log(len(fifteen))
		t.Log("Failure format 2")
		t.Errorf("Any(%v) = false, want true", fifteen)
	}

	// Check for all other style card numbers
	sixteen = "9111222233334444"
	if CreditCard(sixteen).CardIs("AmericanExpress") {
		t.Errorf("Any(%v) = true, want false", sixteen)
	}
}
