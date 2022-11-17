package double

import "errors"

var (
	ErrMissingArg    = errors.New("FirstName and LastName are required")
	ErrNoPersonFound = errors.New("no person found")
)

type Queryer interface {
	Search(people []*Person, firstname, lastname string) *Person
}

type Persons []*Person

type Person struct {
	Firstname string
	Lastname  string
	Phone     string
}

type PhoneBook struct {
	Persons Persons
}

func (p *PhoneBook) Find(query Queryer, firstName, lastName string) (string, error) {
	if firstName == "" || lastName == "" {
		return "", ErrMissingArg
	}

	person := query.Search(p.Persons, firstName, lastName)
	if person == nil {
		return "", ErrNoPersonFound
	}

	return person.Phone, nil
}
