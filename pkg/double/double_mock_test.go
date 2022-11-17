package double

import "testing"

type MockSearcher struct {
	phone         string
	methodsToCall map[string]bool
}

func (ms *MockSearcher) Search(people []*Person, firstname, lastname string) *Person {
	ms.methodsToCall["Search"] = true
	return &Person{Firstname: firstname, Lastname: lastname, Phone: ms.phone}
}

func (ms *MockSearcher) ExpectToCall(methodName string) {
	if ms.methodsToCall == nil {
		ms.methodsToCall = make(map[string]bool)
	}

	ms.methodsToCall[methodName] = false
}

func (ms *MockSearcher) Verify(t *testing.T) {
	for methodName, wasCalled := range ms.methodsToCall {
		if !wasCalled {
			t.Errorf("expected %s to be called", methodName)
		}
	}
}

func TestFindCallsSearchAndReturnsPersonUsingMock(t *testing.T) {
	fakePhone := "+31 65 222 3333"
	phoneBook := &PhoneBook{}

	ms := &MockSearcher{phone: fakePhone}
	ms.ExpectToCall("Search")

	phone, err := phoneBook.Find(ms, "John", "Doe")
	if err != nil {
		t.Errorf("got %v, want %v", err, nil)
	}

	if phone != fakePhone {
		t.Errorf("got %v, want %v", phone, fakePhone)
	}

	ms.Verify(t)
}
