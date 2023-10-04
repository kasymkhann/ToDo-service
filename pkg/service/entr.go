package service

import (
	"crypto/sha1"
	"fmt"
	user "to-doProjectGo"
	"to-doProjectGo/pkg/repository"
)

const salt = "38fdg09dsfv34"

type EntrService struct {
	r repository.Entering
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
