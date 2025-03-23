package passwords

import "golang.org/x/crypto/bcrypt"

type PasswordVerifier struct{}

func NewPasswordVerifier() *PasswordVerifier {
	return &PasswordVerifier{}
}

func (verifier *PasswordVerifier) VerifyPassword(hashedPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
	return err == nil
}
