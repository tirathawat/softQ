package double

import "testing"

type FakeSearcher struct{}

func (fs FakeSearcher) Search(peoples []*Person, firstname, lastname string) *Person {
	if len(peoples) == 0 {
		return nil
	}

	return peoples[0]
}

func FindCallsSearchAndReturnsEmptyStringForNoPersion(t *testing.T) {
	phoneBook := &PhoneBook{}

	phone, err := phoneBook.Find(FakeSearcher{}, "John", "Doe")
	if err != ErrNoPersonFound {
		t.Errorf("got %v, want %v", err, ErrNoPersonFound)
	}

	if phone != "" {
		t.Errorf("got %v, want %v", phone, "")
	}
}
