package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func removeDirectories(directories []string) {
	for _, directory := range directories {

		if err := os.RemoveAll(directory); err != nil {
			fmt.Println(err)
		}
	}
}

func readConfigFile() {
	configFile := os.Getenv("HOME") + "/.houki.yml"

	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", m["Directory"])
}

func main() {
	readConfigFile()
	directories := []string{"tmp"}
	removeDirectories(directories)

	fmt.Println("Have cleaned")
}
