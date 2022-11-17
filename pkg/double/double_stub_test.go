package double

import "testing"

type StubSearcher struct {
	phone string
}

func (ss StubSearcher) Search(people []*Person, firstname, lastname string) *Person {
	return &Person{Firstname: firstname, Lastname: lastname, Phone: ss.phone}
}

func TestFindReturnsPerson(t *testing.T) {
	fakePhone := "+31 65 222 3333"
	phoneBook := &PhoneBook{}

	phone, err := phoneBook.Find(StubSearcher{fakePhone}, "John", "Doe")
	if err != nil {
		t.Errorf("got %v, want %v", err, nil)
	}

	if phone != fakePhone {
		t.Errorf("got %v, want %v", phone, fakePhone)
	}
}
