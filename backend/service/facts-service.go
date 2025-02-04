package service

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
)

type FactsService interface {
	FindAll() []string
	GetRandomFact() string
}

type factsService struct {
	facts []string
}

func New() FactsService {
	pwd, _ := os.Getwd()
	jsonFile, err := os.Open(pwd + "/assets/facts.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("Successfully Opened facts.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var facts []interface{}
	err = json.Unmarshal(byteValue, &facts)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	factsStrings := make([]string, len(facts))
	for i, fact := range facts {
		if str, ok := fact.(string); ok {
			factsStrings[i] = str
		}
	}
	// Retourner une nouvelle instance de factsService avec les faits
	return &factsService{facts: factsStrings}

}

func (service *factsService) FindAll() []string {
	return service.facts
}

func (service *factsService) GetRandomFact() string {
	return service.facts[rand.Intn(len(service.facts))]
}
