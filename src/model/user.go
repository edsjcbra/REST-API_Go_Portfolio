package model

import (
	"crypto/md5"
	"encoding/hex"
)

type user struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}

type UserGetter interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	EncryptPassword()
	SetID(string)
}

func NewUser(email, password, name string, age int8) UserGetter {
	return &user{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func (u *user) SetID(id string) {
	u.id = id
}

func (u *user) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.password))
	u.password = hex.EncodeToString(hash.Sum(nil))
}

func (u *user) GetID() string {
	return u.id
}

func (u *user) GetEmail() string {
	return u.email
}
func (u *user) GetPassword() string {
	return u.password
}
func (u *user) GetName() string {
	return u.name
}
func (u *user) GetAge() int8 {
	return u.age
}
