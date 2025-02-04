package controller

import "github.com/Nils-Abjorn/cat-facts/service"

type FactsController interface {
	FindAll() []string
	GetRandomFact() string
}

type controller struct {
	service service.FactsService
}

func New(service service.FactsService) FactsController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []string {
	return c.service.FindAll()
}

func (c *controller) GetRandomFact() string {
	return c.service.GetRandomFact()
}
