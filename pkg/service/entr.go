package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	user "to-doProjectGo"
	"to-doProjectGo/pkg/repository"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt      = "wfj4334jnb233rb"
	signInKey = "ds2!323hng#384t3b349fbwe"
	tokenTTL  = 12 * time.Hour
)

type EntrService struct {
	r repository.Entering
}

type tokenClaims struct {
	jwt.StandardClaims
	userId int "user_id"
}

type AuthService struct {
	r repository.Entering
}

func (c *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signInKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenlaims")
	}
	return claims.userId, nil

}

func EnteringService(r repository.Entering) *EntrService {
	return &EntrService{r: r}
}

func (e *EntrService) CreateUser(user user.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return e.r.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))

}

func (e *EntrService) GenerateTOKEN(username, password string) (string, error) {
	user, err := e.r.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return " ", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, &tokenClaims{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
	}, user.Id,
	})

	return token.SignedString([]byte(signInKey))
}
