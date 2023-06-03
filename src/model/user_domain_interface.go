package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	SetID(string)

	GetJSONValue() (string, error)
	EncryptPassowrd()
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Name:     name,
		Age:      age,
		Password: password,
	}
}
