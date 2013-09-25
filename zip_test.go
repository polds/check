package check

import (
	"testing"
)

// Hungary
func TestHungaryZip(t *testing.T) {
	var zip string

	zip = "1013"
	if !Zip(zip).OfCountry("hu") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Japan
func TestJapanZip(t *testing.T) {
	var zip string

	zip = "107-0061"
	if !Zip(zip).OfCountry("jp") {
		t.Errorf("%v = false, want true", zip)
	}
}

// United States
func TestUSZip(t *testing.T) {
	var zip string = "83406"

	if !Zip(zip).OfCountry("us") {
		t.Errorf("%v = false, want true", zip)
	}

	zip = "83406-6715"
	if !Zip(zip).OfCountry("us") {
		t.Errorf("%v = false, want true", zip)
	}

}
