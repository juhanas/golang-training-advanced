package secret

var secretKey = "secretKey"

type StringOrInt interface {
	string | int
}

// Common interface that allows encrypting and decrypting data.
// This is useful when we want to be able to call the same functions
// for different types of data/structs, or if we want to have the same
// kind of interface for different things, such as databases.
// By using common interfaces, it will be much easier to change parts of the
// service, such as logging, databases, etc.
type Secreter[T StringOrInt] interface {
	GetName() string
	Encrypt(T) (string, error)
	Decrypt() (T, error)
}
