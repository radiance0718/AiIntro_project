package initialize

import (
	_ "Demo/docs"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var configFile = flag.String("config", "config.json", "Configuration file for server. Defaults to config.json.")

var Configuration Config

type Config struct {
	// Admin information
	AdminName         string                    `json:"AdminName"`
	AdminPassword     string                    `json:"AdminPassWord"`
	CountByType       map[string]int            `json:"CountByType"`
	CountByGrade      map[string]map[string]int `json:"CountByGrade"`
	StudentCategorise map[int]string            `json:"StudentCategorise"`
	StudentImage      map[int]string
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

	image1Info := "./resources/image/Diligent.png"
	image2Info := "./resources/image/Explorer.png"
	image3Info := "./resources/image/Learner.png"
	image4Info := "./resources/image/Researcher.png"
	image5Info := "./resources/image/Thinker.png"
	image6Info := "./resources/image/Unknown.png"
	image7Info := "./resources/image/white.png"

	image1Data, err := ioutil.ReadFile(image1Info)
	if err != nil {
		log.Fatal(err)
	}
	image2Data, err := ioutil.ReadFile(image2Info)
	if err != nil {
		log.Fatal(err)
	}
	image3Data, err := ioutil.ReadFile(image3Info)
	if err != nil {
		log.Fatal(err)
	}
	image4Data, err := ioutil.ReadFile(image4Info)
	if err != nil {
		log.Fatal(err)
	}
	image5Data, err := ioutil.ReadFile(image5Info)
	if err != nil {
		log.Fatal(err)
	}
	image6Data, err := ioutil.ReadFile(image6Info)
	if err != nil {
		log.Fatal(err)
	}
	image7Data, err := ioutil.ReadFile(image7Info)
	if err != nil {
		log.Fatal(err)
	}

	image1Base64 := base64.StdEncoding.EncodeToString(image1Data)
	image2Base64 := base64.StdEncoding.EncodeToString(image2Data)
	image3Base64 := base64.StdEncoding.EncodeToString(image3Data)
	image4Base64 := base64.StdEncoding.EncodeToString(image4Data)
	image5Base64 := base64.StdEncoding.EncodeToString(image5Data)
	image6Base64 := base64.StdEncoding.EncodeToString(image6Data)
	image7Base64 := base64.StdEncoding.EncodeToString(image7Data)

	fullConfig.StudentImage[0] = image1Base64
	fullConfig.StudentImage[1] = image2Base64
	fullConfig.StudentImage[2] = image3Base64
	fullConfig.StudentImage[3] = image4Base64
	fullConfig.StudentImage[4] = image5Base64
	fullConfig.StudentImage[5] = image6Base64
	fullConfig.StudentImage[6] = image7Base64

	Configuration = fullConfig
}
