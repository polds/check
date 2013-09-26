// Package check is a validation library written in Go
// to validate many real world applications.
//
// Examples include: emails, zip codes, phone numbers, etc.
package check

import (
	"regexp"
)

type Email string

// IsValid implements the Email type and
// will validate if the value stored in Email is
// a valid RFC compliant email address.
// This returns a boolean value.
func (e Email) IsValid() bool {

	// BUG(polds) The email regex has issues
	// validating the following types of emails (they return false):
	// user@[IPv6:2001:db8:1ff::a0b:dbd0], postbox@com, and admin@mailserver1 .
	// Which according to RFC 5322 should be valid addresses.
	var pattern string = "^(([^<>()[\\]\\.,;:\\s@\"]+(\\.[^<>()[\\]\\.,;:\\s@\"]+)*)|(\".+\"))@((\\[[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\])|(([a-zA-Z\\-0-9]+\\.)+[a-zA-Z]{2,}))$"

	if string(e) == "" {
		return false
	}

	matched, err := regexp.MatchString(pattern, string(e))
	if err != nil {
		return false
	}

	return matched
}
