package services

import (
	"errors"
	"addressbook.cts.com/models"
)

var contacts []models.Contact

func init() {
	//the init() function in a package is executed once when the package is accessed for the first time.
	//initialize some hypothetical data
	contacts = []models.Contact{}
	contacts = append(contacts, models.Contact{101, "Vamsy", "Kiran", "9052224753", "9550204753", "v.k@iiht.com"})
	contacts = append(contacts, models.Contact{102, "Sagar", "Guru", "9948016004", "9948016664", "s@gmail.com"})
	contacts = append(contacts, models.Contact{103, "Suseela", "Aripaka", "9052444753", "9550244753", "ss@gmail.com"})
}

func indexOf(cid int) int {
	foundAt := -1
	for index, ct := range contacts {
		if ct.ContactId == cid {
			foundAt = index
			break
		}
	}
	return foundAt
}

func GetAllContacts() []models.Contact {
	return contacts
}

func GetContactById(cid int) (models.Contact, error) {

	foundAt := indexOf(cid)

	if foundAt == -1 {
		return models.Contact{}, errors.New("element not found")
	} else {
		return contacts[foundAt], nil
	}
}

func AddContact(contact models.Contact) {
	contacts = append(contacts, contact)
}

func UpdateContact(contact models.Contact) error {
	index := indexOf(contact.ContactId)
	if index == -1 {
		return errors.New("No Such Record")
	} else {
		contacts[index] = contact
		return nil
	}
}

func DeleteContact(cid int) error {
	index := indexOf(cid)
	if index == -1 {
		return errors.New("No Such Record")
	} else {
		copy(contacts[index:], contacts[index+1:])
		contacts = contacts[:len(contacts)-1]
		return nil
	}
}
