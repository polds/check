package check

import (
	"testing"
)

func TestValids(t *testing.T) {
	var validemail string

	validemail = "niceandsimple@example.com"
	if !Email(validemail).IsValid() {
		t.Errorf("%v = false, want true", validemail)
	}

	validemail = "very.common@example.com"
	if !Email(validemail).IsValid() {
		t.Errorf("%v = false, want true", validemail)
	}

	validemail = "a.little.lengthy.but.fine@dept.example.com"
	if !Email(validemail).IsValid() {
		t.Errorf("%v = false, want true", validemail)
	}

	validemail = "disposable.style.email.with+symbol@example.com"
	if !Email(validemail).IsValid() {
		t.Errorf("%v = false, want true", validemail)
	}

	// BUG(polds) Current Regex does not pass
	// validemail = "user@[IPv6:2001:db8:1ff::a0b:dbd0]"
	// if !Email(validemail).IsValid() {
	// 	t.Errorf("%v = false, want true", validemail)
	// }

	validemail = "\"much.more unusual\"@example.com"
	if !Email(validemail).IsValid() {
		t.Errorf("%v = false, want true", validemail)
	}

	validemail = "\"very.unusual.@.unusual.com\"@example.com"
	if !Email(validemail).IsValid() {
		t.Errorf("%v = false, want true", validemail)
	}

	validemail = "\"very.(),:;<>[]\\\".VERY.\\\"very@\\ \\\"very\\\".unusual\"@strange.example.com"
	if !Email(validemail).IsValid() {
		t.Errorf("%v = false, want true", validemail)
	}

	// BUG(polds) Current Regex does not pass
	// validemail = "postbox@com"
	// if !Email(validemail).IsValid() {
	// 	t.Errorf("%v = false, want true", validemail)
	// }

	// BUG(polds) Current Regex does not pass
	// validemail = "admin@mailserver1"
	// if !Email(validemail).IsValid() {
	// 	t.Errorf("%v = false, want true", validemail)
	// }

	validemail = "!#$%&'*+-/=?^_`{}|~@example.org"
	if !Email(validemail).IsValid() {
		t.Errorf("%v = false, want true", validemail)
	}

	validemail = "\"()<>[]:,;@\\\\\"!#$%&'*+-/=?^_`{}| ~.a\"@example.org"
	if !Email(validemail).IsValid() {
		t.Errorf("%v = false, want true", validemail)
	}

	validemail = "\" \"@example.org"
	if !Email(validemail).IsValid() {
		t.Errorf("%v = false, want true", validemail)
	}

	validemail = "üñîçøðé@example.com"
	if !Email(validemail).IsValid() {
		t.Errorf("%v = false, want true", validemail)
	}

}

func TestInvalids(t *testing.T) {
	var invalidemail string

	invalidemail = "Abc.example.com"
	if Email(invalidemail).IsValid() {
		t.Errorf("%v = true, want false", invalidemail)
	}

	invalidemail = "A@b@c@example.com"
	if Email(invalidemail).IsValid() {
		t.Errorf("%v = true, want false", invalidemail)
	}

	invalidemail = "a\"b(c)d,e:f;g<h>i[j\\k]l@example.com"
	if Email(invalidemail).IsValid() {
		t.Errorf("%v = true, want false", invalidemail)
	}

	invalidemail = "just\"not\"right@example.com"
	if Email(invalidemail).IsValid() {
		t.Errorf("%v = true, want false", invalidemail)
	}

	invalidemail = "this is\"not\\allowed@example.com"
	if Email(invalidemail).IsValid() {
		t.Errorf("%v = true, want false", invalidemail)
	}

	invalidemail = "this\\ still\\\"not\\allowed@example.com"
	if Email(invalidemail).IsValid() {
		t.Errorf("%v = true, want false", invalidemail)
	}
}
