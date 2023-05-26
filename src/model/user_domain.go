package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	getEmail() string
	getPassword() string
	getName() string
	getAge() int8

	EncryptPassowrd()
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		email, password, name, age,
	}
}

type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) getEmail() string {
	return ud.email
}
func (ud *userDomain) getPassword() string {
	return ud.password
}
func (ud *userDomain) getName() string {
	return ud.name
}
func (ud *userDomain) getAge() int8 {
	return ud.age
}

func (ud *userDomain) EncryptPassowrd() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
