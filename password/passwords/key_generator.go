package passwords

import "crypto/rand"

const KEY_LEN = 32

type KeyGenerator struct {
	key [KEY_LEN]byte
}

func NewKeyGenerator() (*KeyGenerator, error) {
	var tempKey [KEY_LEN]byte
	_, err := rand.Read(tempKey[:])
	if err != nil {
		return nil, err
	}
	return &KeyGenerator{key: tempKey}, nil
}

func (generator *KeyGenerator) GetKey() [KEY_LEN]byte {
	return generator.key
}
