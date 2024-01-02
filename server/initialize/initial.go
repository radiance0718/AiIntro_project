package initialize

import (
	_ "Demo/docs"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var configFile = flag.String("config", "config.json", "Configuration file for server. Defaults to config.json.")

var Configuration Config

type Config struct {
	// Admin information
	AdminName     string                    `json:"AdminName"`
	AdminPassword string                    `json:"AdminPassWord"`
	CountByType   map[string]int            `json:"CountByType"`
	CountByGrade  map[string]map[string]int `json:"CountByGrade"`
}

func InitConfig() {
	// Read config file
	file, err := os.Open(*configFile)
	if err != nil {
		fmt.Println("Error while trying to read config file: ", err)
		os.Exit(0)
	}
	defer file.Close()

	// Decode config file
	var fullConfig Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&fullConfig)
	if err != nil {
		fmt.Println("Error while trying to load config file: ", err)
		os.Exit(0)
	}
	Configuration = fullConfig
}
