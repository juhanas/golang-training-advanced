package secreter

import (
	"fmt"
)

// var secretKey = "secretKey"

// Common interface that allows encrypting and decrypting data.
// This is useful when we want to be able to call the same functions
// for different types of data/structs, or if we want to have the same
// kind of interface for different things, such as databases.
// By using common interfaces, it will be much easier to change parts of the
// service, such as logging, databases, etc.
type Secreter interface {
	Encrypt(string) (string, error)
	Decrypt() (string, error)
}

// Holds encrypted data and implements the Secreter interface.
// Secreter interface is implemented simply by adding all the methods
// required by the interface (Encrypt & Decrypt).
type Secret struct {
	Name string
	// not exported to allow/suggest changing only via the Encrypt function
	value []byte
}

// Creates and returns a new Secret struct with the given name.
// The struct does not yet contain an actual secret.
func NewSecret(name string) *Secret {
	s := Secret{
		Name: name,
	}
	return &s
}

// Encrypts the given data, stores it in the secret and returns it.
// Returns an error if a secret already exists.
func (s *Secret) Encrypt(data string) (string, error) {
	if len(s.value) != 0 {
		return "", fmt.Errorf("encrypted data already exists")
	}

	s.value = []byte(data)
	return data, nil
}

// Decrypts the data in the struct and returns it.
// Returns error if no data is found.
func (s *Secret) Decrypt() (string, error) {
	if len(s.value) == 0 {
		return "", fmt.Errorf("no secret found")
	}
	return string(s.value), nil
}

// Defining another struct just to showcase interfaces.
// Because the Encrypt & Decrypt methods are defined, also
// this struct implements the Secreter interface.
// The use case for this could be to implement storing
// more complex data than just int or string, such as
// arrays, maps or structs.
type AnotherSecret struct{}

func (a *AnotherSecret) Encrypt(data string) (string, error) {
	return "", nil
}

func (a *AnotherSecret) Decrypt() (string, error) {
	return "", nil
}
