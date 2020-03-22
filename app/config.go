package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type (
	DBConfig struct {
		Database Database
	}

	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		DBName   string `json:"databse"`
		Driver   string `json:"driver"`
	}
)

func (c *DBConfig) GetConfig() {
	configFile, err := os.Open("./config/config.json")
	if err != nil {
		fmt.Printf("Error open the config, %s \n", err.Error())
	}

	byteFIle, err := ioutil.ReadAll(configFile)
	if err != nil {
		fmt.Printf("Error read the config, %s \n", err.Error())
	}

	err = json.Unmarshal(byteFIle, &c)
	if err != nil {
		fmt.Printf("Error marshal the config, %s \n", err.Error())
	}
}
