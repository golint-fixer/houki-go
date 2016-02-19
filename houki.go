package main

import (
	"./prompt"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	Directory []string `yaml:"Directory"`
}

func removeDirectories(directories []string) {
	if ok := prompt.Confirm("Directories\n%s\n\nAre you sure you want to delete directories? ", strings.Join(directories, "\n")); !ok {
		println("Do nothing")
		return
	}

	for _, directory := range directories {
		if err := os.RemoveAll(directory); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Have cleaned")
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

}
