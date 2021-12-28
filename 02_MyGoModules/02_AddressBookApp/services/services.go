package services

import (
	"addressbook.cts.com/models"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const GET_ALL_QRY string = "SELECT * FROM contacts"
const GET_BY_ID_QRY string = "SELECT * FROM contacts WHERE ContactId=?"
const INS_QRY string = "INSERT INTO contacts(FirstName,LastName,Mobile,AlternateMobile,MailId) VALUES(?,?,?,?,?)"
const DEL_QRY string = "DELETE FROM contacts WHERE ContactId=?"

func getConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/contactsdb")

	if err != nil {
		panic(err)
	}

	return db
}

func GetAllContacts() []models.Contact {
	var contacts []models.Contact

	db := getConnection()

	defer db.Close()

	rows, err := db.Query(GET_ALL_QRY)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var contact models.Contact

		err := rows.Scan(&contact.ContactId, &contact.FirstName, &contact.LastName, &contact.Mobile, &contact.AlternateMobile, &contact.MailId)

		if err != nil {
			panic(err)
		}

		contacts = append(contacts, contact)
	}

	return contacts
}

func GetContactById(cid int) (models.Contact, error) {
	var contact models.Contact

	db := getConnection()

	defer db.Close()

	row, err := db.QueryRow(GET_BY_ID_QRY, cid)

	if err == nil {
		err = row.Scan(&contact.ContactId, &contact.FirstName, &contact.LastName, &contact.Mobile, &contact.AlternateMobile, &contact.MailId)
	}

	return contact, err
}

func AddContact(contact models.Contact) {
	db := getConnection()

	defer db.Close()

	_, err := db.Query(INS_QRY, contact.FirstName, contact.LastName, contact.Mobile, contact.AlternateMobile, contact.MailId)

	if err != nil {
		panic(err)
	}
}

func DeleteContact(cid int) error {
	db := getConnection()

	defer db.Close()

	_, err := db.Query(DEL_QRY, cid)

	if err != nil {
		panic(err)
	}
}
