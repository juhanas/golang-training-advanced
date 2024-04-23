package secret

import (
	"fmt"

	elisacommon "github.com/elisasre/go-common"
)

// Creates and returns a new StringSecret struct with the given name.
// The struct does not yet contain an actual secret.
func NewString(name string) *StringSecret {
	s := StringSecret{
		name: name,
	}
	return &s
}

// Holds encrypted data and implements the Secreter interface.
// Secreter interface is implemented simply by adding all the methods
// required by the interface (GetName, Encrypt & Decrypt).
// To allow the interface to be used easily, all internal data
// must be accessible via functions defined in the interface
// (getters & setters) instead of directly accessing the data with
// for example secret.Name.
type StringSecret struct {
	// Values are not exported to allow/suggest changing them
	// only via the specified functions
	name  string
	value []byte
}

// Returns the name of the secret.
// Each secret should have a unique name to allow
// distinguishing between different secrets.
// Note: How the name is used is up to the implementation
// since the interface only requires the method to be defined.
func (s *StringSecret) GetName() string {
	return s.name
}

// Encrypts the given data, stores it in the secret and returns it.
// Returns an error if a secret already exists.
func (s *StringSecret) Encrypt(data string) (string, error) {
	if len(s.value) != 0 {
		return "", fmt.Errorf("encrypted data already exists")
	}

	encryptedValue, err := elisacommon.Encrypt([]byte(data), secretKey)
	if err != nil {
		return "", err
	}

	s.value = encryptedValue
	return string(encryptedValue), nil
}

// Decrypts the data in the struct and returns it.
// Returns error if no data is found.
func (s *StringSecret) Decrypt() (string, error) {
	if len(s.value) == 0 {
		return "", fmt.Errorf("no secret found")
	}

	data, err := elisacommon.Decrypt(s.value, secretKey)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
