package check

import (
	"testing"
)

var zip string

// Argentina
func TestArgentinaZip(t *testing.T) {
	zip = "B1657"
	if !Zip(zip).OfCountry("ar") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Australia
func TestAustraliaZip(t *testing.T) {
	zip = "2000"
	if !Zip(zip).OfCountry("au") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Austria
func TestAustriaZip(t *testing.T) {
	zip = "1010"
	if !Zip(zip).OfCountry("at") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Belgium
func TestBelgiumZip(t *testing.T) {
	zip = "3840"
	if !Zip(zip).OfCountry("be") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Brazil
func TestBrazilZip(t *testing.T) {
	zip = "00000-000"
	if !Zip(zip).OfCountry("br") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Bulgaria
func TestBulgariaZip(t *testing.T) {
	zip = "5094"
	if !Zip(zip).OfCountry("bg") {
		t.Errorf("%v = false, want true", zip)
	}

}

// Canada
func TestCanadaZip(t *testing.T) {
	zip = "L4C 3V2"
	if !Zip(zip).OfCountry("ca") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Croatia
func TestCroatiaZip(t *testing.T) {
	zip = "10000"
	if !Zip(zip).OfCountry("hr") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Cyprus
func TestCyprusZip(t *testing.T) {
	zip = "8501"
	if !Zip(zip).OfCountry("cy") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Czech Republic
func TestCzechRepublicZip(t *testing.T) {
	zip = "160 00"
	if !Zip(zip).OfCountry("cz") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Denmark
func TestDenmarkZip(t *testing.T) {
	zip = "2750"
	if !Zip(zip).OfCountry("dk") {
		t.Errorf("%v = false, want true", zip)
	}

	zip = "1750"
	if Zip(zip).OfCountry("dk") {
		t.Errorf("%v = false, want true", zip)
	}

}

// Estonia
func TestEstoniaZip(t *testing.T) {
	zip = "42106"
	if !Zip(zip).OfCountry("ee") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Finland
func TestFinlandZip(t *testing.T) {
	zip = "55100"
	if !Zip(zip).OfCountry("fi") {
		t.Errorf("%v = false, want true", zip)
	}
}

// France
func TestFranceZip(t *testing.T) {
	zip = "52110"
	if !Zip(zip).OfCountry("fr") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Germany
func TestGermanyZip(t *testing.T) {
	zip = "79258"
	if !Zip(zip).OfCountry("de") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Great Britain
func TestGreatBritainZip(t *testing.T) {
	zip = "EC1A 1BB"
	if !Zip(zip).OfCountry("gb") {
		t.Errorf("%v = false, want true", zip)
	}

	zip = "W1A 1HQ"
	if !Zip(zip).OfCountry("gb") {
		t.Errorf("%v = false, want true", zip)
	}

	zip = "M1 1AA"
	if !Zip(zip).OfCountry("gb") {
		t.Errorf("%v = false, want true", zip)
	}

	zip = "B33 8TH"
	if !Zip(zip).OfCountry("gb") {
		t.Errorf("%v = false, want true", zip)
	}

	zip = "CR2 6XH"
	if !Zip(zip).OfCountry("gb") {
		t.Errorf("%v = false, want true", zip)
	}

	zip = "DN55 1PT"
	if !Zip(zip).OfCountry("gb") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Greece
func TestGreeceZip(t *testing.T) {
	zip = "681 00"
	if !Zip(zip).OfCountry("gr") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Hungary
func TestHungaryZip(t *testing.T) {
	zip = "1013"
	if !Zip(zip).OfCountry("hu") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Iceland
func TestIcelandZip(t *testing.T) {
	zip = "720"
	if !Zip(zip).OfCountry("is") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Italy
func TestItalyZip(t *testing.T) {
	zip = "26812"
	if !Zip(zip).OfCountry("it") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Japan
func TestJapanZip(t *testing.T) {
	zip = "107-0061"
	if !Zip(zip).OfCountry("jp") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Latvia
func TestLatviaZip(t *testing.T) {
	zip = "LV-3701"
	if !Zip(zip).OfCountry("lv") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Lithuania
func TestLithuaniaZip(t *testing.T) {
	zip = "73461"
	if !Zip(zip).OfCountry("lt") {
		t.Errorf("%v = false, want true", zip)
	}

	zip = "LT-73184"
	if !Zip(zip).OfCountry("lt") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Luxembourg
func TestLuxembourgZip(t *testing.T) {
	// zip = "1010"
	// if !Zip(zip).OfCountry("lu") {
	// 	t.Errorf("%v = false, want true", zip)
	// }

	zip = "L-2920"
	if !Zip(zip).OfCountry("lu") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Malta
func TestMaltaZip(t *testing.T) {
	zip = "BBG 1014"
	if !Zip(zip).OfCountry("mt") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Netherlands
func TestNetherlandsZip(t *testing.T) {
	zip = "1000 AP"
	if !Zip(zip).OfCountry("nl") {
		t.Errorf("%v = false, want true", zip)
	}

	// zip = "1000"
	// if !Zip(zip).OfCountry("nl") {
	// 	t.Errorf("%v = false, want true", zip)
	// }
}

// Norway
func TestNorwayZip(t *testing.T) {
	zip = "0001"
	if !Zip(zip).OfCountry("no") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Poland
func TestPolandZip(t *testing.T) {
	zip = "26-600"
	if !Zip(zip).OfCountry("pl") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Portugal
func TestPortugalZip(t *testing.T) {
	zip = "1050"
	if !Zip(zip).OfCountry("pt") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Romania
func TestRomaniaZip(t *testing.T) {
	zip = "827019"
	if !Zip(zip).OfCountry("ro") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Slovakia
func TestSlovakiaZip(t *testing.T) {
	zip = "811 02"
	if !Zip(zip).OfCountry("sk") {
		t.Errorf("%v = false, want true", zip)
	}

	zip = "SK-811 02"
	if !Zip(zip).OfCountry("sk") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Slovenia
func TestSloveniaZip(t *testing.T) {
	zip = "1233"
	if !Zip(zip).OfCountry("si") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Spain
func TestSpainZip(t *testing.T) {
	zip = "33559"
	if !Zip(zip).OfCountry("es") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Sweden
func TestSwedenZip(t *testing.T) {
	zip = "254 76"
	if !Zip(zip).OfCountry("se") {
		t.Errorf("%v = false, want true", zip)
	}

	zip = "25476"
	if !Zip(zip).OfCountry("se") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Turkey
func TestTurkeyZip(t *testing.T) {
	zip = "21500"
	if !Zip(zip).OfCountry("tr") {
		t.Errorf("%v = false, want true", zip)
	}
}

// United States
func TestUSZip(t *testing.T) {
	zip = "83406"
	if !Zip(zip).OfCountry("us") {
		t.Errorf("%v = false, want true", zip)
	}

	zip = "83406-6715"
	if !Zip(zip).OfCountry("us") {
		t.Errorf("%v = false, want true", zip)
	}
}

// Ukraine
func TestUkraineZip(t *testing.T) {
	zip = "27420"
	if !Zip(zip).OfCountry("ua") {
		t.Errorf("%v = false, want true", zip)
	}
}
