package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/segmentio/go-prompt"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Directory []string `yaml:"Directory"`
}

func reCreateDirectory(directory string) {
	if err := os.RemoveAll(directory); err != nil {
		fmt.Println(err)
	} else {
		os.Mkdir(directory, 0775)
	}
}

func removeDirectories(directories []string) {
	if ok := prompt.Confirm("Directories\n%s\n\nAre you sure you want to delete directories? ", strings.Join(directories, "\n")); !ok {
		return
	}

	for _, directory := range directories {
		go reCreateDirectory(directory)
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
