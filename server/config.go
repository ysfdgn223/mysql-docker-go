package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Config interface {
	GetPort() string
	Print()
}

type config struct {
	Port string
}

func NewConfig() Config {
	byteFile, err := ioutil.ReadFile(".conf/config.json")

	if err != nil {
		log.Fatal("error when reading config.json, err : ", err)
	}

	var jFile jsonFile

	err = json.Unmarshal(byteFile, &jFile)

	if err != nil {
		log.Fatal("error when unmarshalling json bytefile", err)
	}

	return &config{Port: jFile.P}
}

func (c *config) GetPort() string {
	p := fmt.Sprintf(":%s", c.Port)
	return p
}

func (c *config) Print() {
	fmt.Println("====server.Config====")
	fmt.Printf("Port : %s\n", c.Port)
}

type jsonFile struct {
	P string `json:"port"`
}
