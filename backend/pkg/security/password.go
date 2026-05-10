package security

import "golang.org/x/crypto/bcrypt"

type PasswordManager struct{}

func NewPasswordManager() *PasswordManager {
	return &PasswordManager{}
}

func (m *PasswordManager) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (m *PasswordManager) Compare(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
