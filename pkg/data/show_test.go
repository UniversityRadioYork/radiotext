package data

import "testing"

func TestWithInTitle(t *testing.T) {
	if !withInTitle("A Show Name with a Presenter") {
		t.Fail()
	}
}

func TestWithNotInTitle(t *testing.T) {
	if withInTitle("The word is not here in this show name") {
		t.Fail()
	}
}
