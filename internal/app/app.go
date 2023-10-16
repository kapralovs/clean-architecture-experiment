package app

import (
	"clean-arch-exp/internal/controllers"
	"clean-arch-exp/internal/usecase"
	"clean-arch-exp/internal/usecase/repository"
)

func Run() {
	repo := repository.New()

	uc := usecase.New(repo)

	controller := controllers.New(uc)

	controller.Serve()
}
