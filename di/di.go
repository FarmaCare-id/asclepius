package di

import (
	"farmacare/controller"
	"farmacare/interfaces"
	"farmacare/repository"
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

	if err := interfaces.Register(Container); err != nil {
		log.Fatal(err.Error())
	}

	if err := controller.Register(Container); err != nil {
		log.Fatal(err.Error())
	}
}