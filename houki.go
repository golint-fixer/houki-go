package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Directory []string `yaml:"Directory"`
}

func removeDirectories(directories []string) {
	for _, directory := range directories {

		if err := os.RemoveAll(directory); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Remove: %s\n", directory)
	}
}

func readConfigFile() Config {
	configFile := os.Getenv("HOME") + "/.houki.yml"

	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	var parsedMap Config
	if err = yaml.Unmarshal(buf, &parsedMap); err != nil {
		panic(err)
	}

	return parsedMap
}

func main() {
	config := readConfigFile()
	removeDirectories(config.Directory)

	fmt.Println("Have cleaned")
}
