package di

import (
	"farmacare/application"
	"farmacare/infrastructure"
	"farmacare/interfaces"
	"farmacare/shared"
	"log"

	"go.uber.org/dig"
)

var Container = dig.New()

func init() {
	if err := shared.Register(Container); err != nil {
		log.Fatal(err.Error())
	}

	if err := application.Register(Container); err != nil {
		log.Fatal(err.Error())
	}

	if err := interfaces.Register(Container); err != nil {
		log.Fatal(err.Error())
	}

	if err := infrastructure.Register(Container); err != nil {
		log.Fatal(err.Error())
	}
}