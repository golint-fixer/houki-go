package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/apcera/termtables"
	"github.com/segmentio/go-prompt"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Directory []string `yaml:"Directory"`
}

func reCreateDirectory(directory string, wg *sync.WaitGroup) {
	defer wg.Done()
	if err := os.RemoveAll(directory); err != nil {
		fmt.Println(err)
	} else {
		os.Mkdir(directory, 0775)
	}
}

func removeDirectories(directories []string) {
	table := termtables.CreateTable()
	table.AddHeaders("Directories")
	for _, directory := range directories {
		table.AddRow(directory)
	}

	if ok := prompt.Confirm("%s\nAre you sure you want to delete directories? ", table.Render()); !ok {
		return
	}

	var wg sync.WaitGroup
	for _, directory := range directories {
		wg.Add(1)
		go reCreateDirectory(directory, &wg)
	}
	wg.Wait()
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
