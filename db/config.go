package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Config interface {
	GetDBConnectionString() (string, string)
	Print()
}

type config struct {
	DbmsName string `json:"dbms_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
}

func NewDBConfig() Config {
	bytef, err := ioutil.ReadFile(".conf/db.json")

	if err != nil {
		log.Fatal("error reading file db.json, err : ", err)
	}

	var d config

	err = json.Unmarshal(bytef, &d)

	if err != nil {
		log.Fatal("error unmarshalling file db.json, err : ", err)
	}

	return &d
}

func (c *config) GetDBConnectionString() (string, string) {
	dbmsName := c.DbmsName
	str := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Username, c.Password, c.Host, c.Port, c.DbName)

	return dbmsName, str
}

func (c *config) Print() {
	fmt.Println("====db.Config====")
	fmt.Printf(
		"DbmsName : %s \nUsername : %s \nPassword : %s \nHost : %s \nPort : %s \nDbName : %s\n", c.DbmsName, c.Username, c.Password, c.Host, c.Port, c.DbName,
	)
}
