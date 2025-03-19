package models

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CacheStore map[string]User

type Conn struct {
	store CacheStore
}

func NewConn() *Conn {
	return &Conn{store: make(CacheStore, 100)}
}

func (c *Conn) CreateUser(n NewUser) (User, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(n.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	_, ok := c.store[n.Email]
	if ok {
		return User{}, errors.New("user already exists")
	}
	us := User{
		Id:           uuid.NewString(),
		Email:        n.Email,
		Name:         n.Name,
		Age:          n.Age,
		PasswordHash: string(passHash),
	}
	c.store = CacheStore{us.Email: us}
	return us, nil

}
