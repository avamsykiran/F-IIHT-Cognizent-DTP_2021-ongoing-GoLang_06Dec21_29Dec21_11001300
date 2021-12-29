DROP DATABASE contactsdb;

CREATE DATABASE contactsdb;

USE contactsdb;

CREATE TABLE contacts(
    ContactId       INT AUTO_INCREMENT,
	FirstName       VARCHAR(50) NOT NULL,
	LastName        VARCHAR(50) NOT NULL,
	Mobile          CHAR(10) NOT NULL,
	AlternateMobile CHAR(10) NOT NULL,
	MailId          VARCHAR(50) NOT NULL,
    PRIMARY KEY (ContactId)
);

INSERT INTO contacts(FirstName,LastName,Mobile,AlternateMobile,MailId)
VALUES('Vamsy','Aripaka','9052224753','9550204753','vamsy.kiran@iiht.com');

INSERT INTO contacts(FirstName,LastName,Mobile,AlternateMobile,MailId)
VALUES('Suseela','Aripaka','9948016004','9999988888','suseela@gmail.com');

INSERT INTO contacts(FirstName,LastName,Mobile,AlternateMobile,MailId)
VALUES('Indhikaa','Aripaka','8787878787','7878787878','indhikaa@okridge.com');

INSERT INTO contacts(FirstName,LastName,Mobile,AlternateMobile,MailId)
VALUES('Srinu','Dachepalli','9898989898','8989898989','srinu@gmail.com');
