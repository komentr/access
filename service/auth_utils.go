package service

import (
	"errors"
	"log"
	"os"
	"regexp"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"gitlab.com/komentr/access/models"
	"golang.org/x/crypto/bcrypt"
)

//DefaultPass const
const (
	DefaultPass       = "DEFAULTPASS"
	MinPassLength     = 6
	MinUsernameLength = 5
	MaxUsernameLength = 40
)

func getDefaultPass() string {
	secret := os.Getenv(DefaultPass)
	if secret == "" {
		secret = DefaultPass
	}
	return secret
}

//ValidateRegistration func
func ValidateRegistration(req models.UserRequest) error {
	if err := ValidateUsername(req.Username); err != nil {
		return err
	}

	if err := ValidatePassword(req.Password); err != nil {
		return err
	}

	return nil
}

//ValidateUsername func
func ValidateUsername(username string) error {
	regexUsername := regexp.MustCompile("^[a-z0-9_\\.-]*$")

	if len(username) < MinUsernameLength || len(username) > MaxUsernameLength {
		return errors.New("username length must be between 5 and 40 characters")
	}

	if !regexUsername.MatchString(username) {
		return errors.New("username format invalid")
	}

	return nil
}

func validateEmail(email string) error {
	rEmail := regexp.MustCompile("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$")

	if !rEmail.MatchString(email) {
		return errors.New("email format invalid")
	}

	return nil
}

//ValidatePassword func
func ValidatePassword(password string) error {
	if len(password) < MinPassLength {
		return errors.New("password length less than 6 characters")
	}
	return nil

}

//HashPassword func
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash func
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//GenerateToken func
func GenerateToken(signingKey []byte, userid, tid string, role, exp int) (string, error) {
	var claims models.CustomClaims
	if exp > 0 {
		claims = models.CustomClaims{
			userid,
			tid,
			role,
			jwt.StandardClaims{
				// ExpiresAt: time.Now().Add(time.Second * expiration).Unix(),
				ExpiresAt: time.Now().Add(time.Hour * time.Duration(exp)).Unix(),
				IssuedAt:  jwt.TimeFunc().Unix(),
			},
		}
	} else {
		claims = models.CustomClaims{
			userid,
			tid,
			role,
			jwt.StandardClaims{
				IssuedAt: jwt.TimeFunc().Unix(),
			},
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

//Register func
func Register(req models.UserRequest) (models.User, error) {
	s := models.User{}

	temppassword := getDefaultPass()
	if req.Password != "" {
		temppassword = req.Password
	}
	hashedPass, err := HashPassword(temppassword)
	if err != nil {
		log.Println("Registration failed, reason=", err)
		return s, err
	}

	s.Username = req.Username
	s.Password = hashedPass
	s.Name = req.Name
	s.Role = req.Role
	s.Image = req.Image
	s.Active = req.Active
	s.RegBy = ""
	s.TID = ""
	s.Created = time.Now()

	return s, nil
}
