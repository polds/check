package check

import (
	"regexp"
)

type Email string

func (e Email) IsValid() bool {
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
