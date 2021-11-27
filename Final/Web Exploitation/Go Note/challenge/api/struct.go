package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

type User struct {
	ID       int64
	UserName string
	Password string
	Email    string
	Data     []Note
}

type Note struct {
	ID      int64
	Data    string
	Title   string
	UsersID int64
	Encrypt bool
}

func (b *User) FixUserName() *User {
	b.UserName = strings.ToLower(b.UserName)
	return b
}

func (b *User) Add() (*User, error) {
	var idusers int64
	row := dbsqlite.QueryRow("Select id from Users where username=?", b.UserName)
	err := row.Scan(&idusers)
	if err != nil && err == sql.ErrNoRows {
		stmt, err := dbsqlite.Prepare("INSERT INTO Users(username,password,email) values(?,?,?)")
		if err != nil {
			return nil, err
		}

		res, err := stmt.Exec(b.UserName, b.Password, b.Email)
		if err != nil {
			return nil, err
		}

		idusers, err = res.LastInsertId()
		if err != nil {
			return nil, err
		}

		b.Password = ""
		b.AddUserID(idusers)

		return b, nil
	} else {
		b.AddUserID(idusers)
		return b, errors.New("username already taken")
	}
}

func (b *User) AddUserID(id int64) *User {
	b.ID = id
	return b
}

func (b *User) Login() (*User, error) {
	var tmpUser User
	row := dbsqlite.QueryRow("select * from Users where username=?", b.UserName)
	err := row.Scan(&tmpUser.ID, &tmpUser.UserName, &tmpUser.Password, &tmpUser.Email)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else {
		if b.Password != tmpUser.Password {
			return nil, errors.New("invalid password")
		}
		return &tmpUser, nil
	}
}

func (b *User) BcryptPass() *User {
	sha256 := sha256.Sum256([]byte(b.Password))
	b.Password = string(sha256[:])
	return b
}

func (b *User) GetNotes() ([]Note, error) {
	rows, err := dbsqlite.Query("SELECT * FROM Note where UsersID=?", b.ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var NoteTMP Note
		err = rows.Scan(&NoteTMP.ID, &NoteTMP.Data, &NoteTMP.UsersID, &NoteTMP.Encrypt, &NoteTMP.Title)
		if err != nil {
			return nil, err
		}

		if NoteTMP.Encrypt {
			NoteTMP.Data = decrypt(NoteTMP.Data, NoteTMP.UsersID)
		}

		b.Data = append(b.Data, NoteTMP)
	}
	return b.Data, nil

}

func (b *Note) AddNote() error {
	stmt, err := dbsqlite.Prepare("INSERT INTO Note(NoteUser,UsersID,Encrypt,Title) values(?,?,?,?)")
	if err != nil {
		return err
	}

	if b.Encrypt {
		_, err = stmt.Exec(encrypt(b.Data, b.UsersID), b.UsersID, b.Encrypt, b.Title)
		if err != nil {
			return err
		}
	} else {
		_, err = stmt.Exec(b.Data, b.UsersID, b.Encrypt, b.Title)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Note) UpdateNote() error {
	stmt, err := dbsqlite.Prepare("UPDATE Note Set NoteUser=?,Title=? where UsersID=? and ID=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(b.Data, b.Title, b.UsersID, b.ID)
	if err != nil {
		return err
	}

	return nil
}

func (b *Note) DeleteNote() error {
	stmt, err := dbsqlite.Prepare("Delete From Note where UsersID=? and ID=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(b.UsersID, b.ID)
	if err != nil {
		return err
	}

	return nil
}

func getKey(id int64) cipher.Block {
	oriKey := []byte(os.Getenv("KEY"))
	oriKey[len(oriKey)-1] = byte(id)
	c, err := aes.NewCipher(oriKey)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func encrypt(stringToEncrypt string, id int64) (encryptedString string) {
	plaintext := []byte(stringToEncrypt)

	aesGCM, err := cipher.NewGCM(getKey(id))
	if err != nil {
		log.Fatal(err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

func decrypt(encryptedString string, id int64) (decryptedString string) {
	enc, _ := hex.DecodeString(encryptedString)
	aesGCM, err := cipher.NewGCM(getKey(id))
	if err != nil {
		log.Fatal(err)
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}

var lipsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec hendrerit malesuada erat placerat convallis.`
