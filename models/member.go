package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
)

type Member struct {
	email     string
	id        int
	password  string
	firstName string
}

func (this *Member) Email() string {
	return this.email
}

func (this *Member) Id() int {
	return this.id
}

func (this *Member) Password() string {
	return this.password
}

func (this *Member) FirstName() string {
	return this.firstName
}

func (this *Member) SetEmail(value string) {
	this.email = value
}

func (this *Member) SetId(value int) {
	this.id = value
}

func (this *Member) SetPassword(value string) {
	this.password = value
}

func (this *Member) SetFirstName(value string) {
	this.firstName = value
}

func GetMember(email string, password string) (Member, error) {
	db, err := getDBConnection()
	if err == nil {
		defer db.Close()
		pwd := sha256.Sum256([]byte(password))
		row := db.QueryRow(`select id, email, firstName from Memeber where email = $1 AND password = $2`, email, hex.EncodeToString(pwd[:]))
		result := Member{}
		err = row.Scan(&result.id, &result.email, &result.firstName)
		if err == nil {
			return result, nil
		} else {
			return result, errors.New("Unable to find member with email: " + email)
		}
	} else {
		return Member{}, errors.New("Unable to get database connection")
	}

}

type Session struct {
	id        int
	memberId  int
	sessionId string
}

func (this *Session) Id() int {
	return this.id
}

func (this *Session) MemberId() int {
	return this.memberId
}

func (this *Session) SessionId() string {
	return this.sessionId
}

func (this *Session) SetId(value int) {
	this.id = value
}

func (this *Session) SetMemberId(value int) {
	this.memberId = value
}

func (this *Session) SetSessionId(value string) {
	this.sessionId = value
}

func CreateSession(member Member) (Session, error) {
	result := Session{}
	result.memberId = member.Id()
	sessionId := sha256.Sum256([]byte(member.Email() + time.Now().Format("12:00:00")))
	result.sessionId = hex.EncodeToString(sessionId[:])
	db, err := getDBConnection()
	if err == nil {
		defer db.Close()
		err := db.QueryRow(`Insert into session (member_id, session_id) values ($1, $2) returning id`, member.Id(), result.sessionId).Scan(&result.id)
		if err == nil {
			return result, nil
		} else {
			return Session{}, errors.New("Unable to save seeion to database")
		}
	} else {
		return result, errors.New("Unable to get database Connection")
	}
}
