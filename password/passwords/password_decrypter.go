package passwords

import "github.com/gtank/cryptopasta"

type PasswordDecrypter struct {
	encryptedPassword []byte
	generator         *KeyGenerator
}

func NewPasswrodDecrypter(generator *KeyGenerator, encryptedPassword []byte) *PasswordDecrypter {
	return &PasswordDecrypter{generator: generator, encryptedPassword: encryptedPassword}
}

func (decrypter *PasswordDecrypter) GetDecryptedPassword() (string, error) {
	key := decrypter.generator.GetKey()
	decrypted, err := cryptopasta.Decrypt(decrypter.encryptedPassword, &key)
	if err != nil {
		return "", nil
	}
	return string(decrypted), err

}
