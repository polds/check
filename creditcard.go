package check

import (
	"regexp"
)

type CreditCard string

func (c CreditCard) patterns() map[string]string {
	patterns := map[string]string{
		"Any":             "^[0-9]{15,16}$",
		"AmericanExpress": "^(34)|(37)\\d{14}$",
		"Discover":        "^6011\\d{12}$",
		"Visa":            "^4\\d{15}$",
		"MasterCard":      "^5\\d{15}$",
	}

	return patterns
}

// CardIs will validate if the value of CreditCard
// is of the provided type.
//
// Valid options:
// Any, AmericanExpress, Visa, MasterCard
func (c CreditCard) CardIs(card string) bool {
	if c.patterns()[card] == "" {
		return false
	}

	return c.match(c.patterns()[card], string(c))
}

// Merchant will get the merchant (card type)
// of a provided credit card number.
//
// If no card is found "Not a valid credit card" is returned.
func (c CreditCard) Merchant() string {
	// @todo return string, error. Let user translate error themselves.
	var out string = "Not a valid credit card"
	if !c.CardIs("Any") {
		return out
	}

	for card, pattern := range c.patterns() {
		if card == "Any" {
			continue
		}
		if c.match(pattern, string(c)) {
			out = card
			break
		}
	}

	return out
}

func (c CreditCard) match(pattern string, s string) bool {
	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		return false
	}

	return matched
}
