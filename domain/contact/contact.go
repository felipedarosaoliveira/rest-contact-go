package contact

import (
	"github.com/felipedarosaoliveira/rest-contact-go/domain"
)

type service struct {
	contacts []domain.Contact
	count    int64
}

//NewService create a new service
func NewService() *service {
	return &service{contacts: []domain.Contact{domain.Contact{ID: 1, Name: "Contact 01", Email: "contact01@test.com", Phone: "33221100"}}, count: 2}
}

func (s *service) LoadAllContacts() []domain.Contact {
	return s.contacts
}
func (s *service) FindContactByID(ID int64) domain.Contact {
	contact := domain.Contact{}
	for _, value := range s.contacts {
		if ID == value.ID {
			contact = value
			break
		}
	}
	return contact
}
func (s *service) Insert(contact *domain.Contact) int64 {
	contact.ID = s.count
	s.count++
	s.contacts = append(s.contacts, *contact)
	return contact.ID
}
func (s *service) Update(contact *domain.Contact) bool {
	result := false
	for index, value := range s.contacts {
		if value.ID == contact.ID {
			s.contacts[index] = domain.Contact{ID: contact.ID, Name: contact.Name, Email: contact.Email, Phone: contact.Phone}
			result = true
			break
		}
	}
	return result
}
func (s *service) Remove(ID int64) bool {
	result := false
	for index, value := range s.contacts {
		if value.ID == ID {
			s.contacts = append(s.contacts[:index], s.contacts[index+1:]...)
			result = true
			break
		}
	}
	return result
}
