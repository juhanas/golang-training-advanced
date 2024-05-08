package secret

// Creates and returns a new IntSecret struct with the given name.
// The struct does not yet contain an actual secret.
func NewNumber(name string) *NumberSecret {
	s := NumberSecret{
		name: name,
	}
	return &s
}

// Defining another struct just to showcase interfaces.
// Because the GetName, Encrypt & Decrypt methods are defined,
// also this struct implements the Secreter interface.
// The use case for this could be to implement storing
// more complex data than just int or string, such as
// arrays, maps or structs (but they need to be added
// using the string data field).
type NumberSecret struct {
	name string
}

func (a *NumberSecret) GetName() string {
	return ""
}

func (a *NumberSecret) Encrypt(data string) (string, error) {
	return "", nil
}

func (a *NumberSecret) Decrypt() (string, error) {
	return "", nil
}
