package authentication

import (
	"errors"
	"server/models"

	"github.com/lucas11776-golang/orm"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid user credentials")
)

// Create new account in users table
func Register(form map[string]string) (*models.User, error) {
	hash, err := HashPassword(form["password"])

	if err != nil {
		return nil, err
	}

	return orm.Model(models.User{}).Insert(orm.Values{
		"first_name": form["first_name"],
		"last_name":  form["last_name"],
		"email":      form["email"],
		"password":   hash,
	})
}

// Check if email and password match record in users table
func Login(email string, password string) (*models.User, error) {
	user, err := orm.Model(models.User{}).
		Where("email", "=", email).
		First()

	if err != nil || user == nil || !CheckHash(password, user.Password) {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

// Hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

// Check if has match password
func CheckHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
