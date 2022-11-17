package double

import "testing"

type SpySearcher struct {
	phone           string
	searchWasCalled bool
}

func (ss *SpySearcher) Search(people []*Person, firstname, lastname string) *Person {
	ss.searchWasCalled = true
	return &Person{Firstname: firstname, Lastname: lastname, Phone: ss.phone}
}

func TestFindCallsSearchAndReturnsPerson(t *testing.T) {
	fakePhone := "+31 65 222 3333"
	phoneBook := &PhoneBook{}

	ss := &SpySearcher{phone: fakePhone}
	phone, err := phoneBook.Find(ss, "John", "Doe")
	if err != nil {
		t.Errorf("got %v, want %v", err, nil)
	}

	if phone != fakePhone {
		t.Errorf("got %v, want %v", phone, fakePhone)
	}

	if !ss.searchWasCalled {
		t.Error("search was not called")
	}
}
