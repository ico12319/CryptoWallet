package passwords

import (
	"github.com/gtank/cryptopasta"
)

type PasswordEncrypter struct {
	generator *KeyGenerator
}

func NewPasswordEncrypter() (*PasswordEncrypter, error) {
	dummyGenerator, err := NewKeyGenerator()
	if err != nil {
		return nil, err
	}
	return &PasswordEncrypter{generator: dummyGenerator}, nil
}

func (encypter *PasswordEncrypter) EncryptPassword(password string) ([]byte, error) {
	key := encypter.generator.GetKey()
	encryptedPassword, err := cryptopasta.Encrypt([]byte(password), &key)
	if err != nil {
		return nil, err
	}
	return encryptedPassword, nil
}

func (encrypter *PasswordEncrypter) GetGenerator() *KeyGenerator {
	return encrypter.generator
}
