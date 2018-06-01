package domain

//this file define the public interface of comunication about contact

//ContactService define actions  about Contact
type ContactService interface {
	LoadAllContacts() []Contact
	FindContactByID(ID int64) Contact
	Insert(contact *Contact) int64
	Update(contact *Contact) bool
	Remove(ID int64) bool
}

//Contact is a simple contact
type Contact struct {
	ID    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

func (c *Contact) String() string {
	return c.Name
}
