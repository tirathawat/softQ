package double

import "testing"

type DummaySearcher struct{}

func (ds DummaySearcher) Search(people []*Person, firstname, lastname string) *Person {
	return &Person{}
}

func TestFindItShouldReturnsErrorWhenFirstNameReturnOrLastNameEmpty(t *testing.T) {
	phonebook := &PhoneBook{}
	want := ErrMissingArg

	_, got := phonebook.Find(DummaySearcher{}, "", "")

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
