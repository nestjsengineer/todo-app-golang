// services/user_service.go
package services

import (
	"todo-app/pkg/models"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) RegisterUser(username, password string) (*models.User, error) {
	// Hash the user's password before storing it
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Password: string(passwordHash),
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(username, password string) (*models.User, error) {
	user := &models.User{}
	if err := s.db.Where("username = ?", username).First(user).Error; err != nil {
		return nil, err
	}

	// Compare the provided password with the stored password hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}

var secretKey = []byte("your-secret-key") // Replace with your secret key

// CreateToken generates a new JWT token.
func (s *UserService) CreateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expiration time (1 hour)
	tokenString, err := token.SignedString(secretKey)
	return tokenString, err
}

// VerifyToken verifies the validity of a JWT token.
func (s *UserService) VerifyToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	return err == nil && token.Valid
}
