package check

import (
	"regexp"
)

type Zip string

func (z Zip) patterns() map[string]string {
	patterns := map[string]string{
		"ar": "^[A-Za-z]\\d{4}$",
		"au": "^\\d{4}$",
		"at": "^\\d{4}$",
		"be": "^\\d{4}$",
		"bg": "^\\d{4}$",
		"br": "^\\d{5}[\\-]?\\d{3}$",
		"ca": "^[A-Za-z]\\d[A-Za-z] \\d[A-Za-z]\\d$",
		"cy": "^\\d{4}$",
		"dk": "^(2)|(3)\\d{2,3}$",
		"de": "^\\d{5}$",
		"es": "^((0[1-9]|5[0-2])|[1-4]\\d)\\d{3}$",
		"gb": "^[A-Za-z]{1,2}[0-9Rr][0-9A-Za-z]? \\d[ABD-HJLNP-UW-Zabd-hjlnp-uw-z]{2}$",
		"hu": "^\\d{4}$",
		"hr": "^\\d{5}$",
		"is": "^\\d{3}$",
		"it": "^\\d{5}$",
		"jp": "^\\d{3}-\\d{4}$",
		"lu": "^(L\\s*(-|—|–))\\s*?[\\d]{4}$",
		"nl": "^[1-9]\\d{3}\\s?[a-zA-Z]{2}$",
		"pl": "^\\d{2}\\-\\d{3}$",
		"se": "^\\d{3}\\s?\\d{2}$",
		"us": "^(\\d{5}([\\-]\\d{4})?)$",
	}

	return patterns
}

func (z Zip) OfCountry(c string) bool {
	if z.patterns()[c] == "" {
		return false
	}

	return z.match(z.patterns()[c], string(z))
}

func (z Zip) match(pattern string, s string) bool {
	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		return false
	}

	return matched
}
