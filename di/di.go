package di

import (
	"farmacare/controller"
	"farmacare/repository"
	"farmacare/service"
	"farmacare/shared"
	"log"

	"go.uber.org/dig"
)

var Container = dig.New()

func init() {
	if err := shared.Register(Container); err != nil {
		log.Fatal(err.Error())
	}

	if err := repository.Register(Container); err != nil {
		log.Fatal(err.Error())
	}

	if err := service.Register(Container); err != nil {
		log.Fatal(err.Error())
	}

	if err := controller.Register(Container); err != nil {
		log.Fatal(err.Error())
	}
}