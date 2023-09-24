package util

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func GetEnv(name string, path string, stru interface{}) {
	viper.SetConfigName(name)   // name of config file (without extension)
	viper.SetConfigType("yml")  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)   // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	err = viper.Unmarshal(&stru)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func PrintEnv(stru interface{}) {
	yamlData, err := yaml.Marshal(&stru)

	if err != nil {
		log.Fatalf("Error while Marshaling. %v", err)
	}

	err = os.WriteFile(".env.example.yml", yamlData, fs.ModeAppend)
	if err != nil {
		log.Fatalf("Error creating file. %v", err)
	}
	fmt.Println("environment file name .env.example.yml is create in base path")
	fmt.Println(string(yamlData))
}
