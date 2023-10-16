package repository

import (
	"clean-arch-exp/internal/models"
	"fmt"
	"sync"
)

type localRepo struct {
	storage map[int]*models.User
	mu      *sync.Mutex
}

func New() *localRepo {
	return &localRepo{
		storage: map[int]*models.User{
			1: &models.User{
				ID:   1,
				Name: "Sergey",
				Age:  27,
			},
		},
		mu: new(sync.Mutex),
	}
}

func (lr *localRepo) Get() {
	fmt.Println(lr.storage[1])
}
