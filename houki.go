package main

import (
	"fmt"
	"os"
)

func removeDirectories(directories []string) {
	for _, directory := range directories {

		if err := os.RemoveAll(directory); err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	directories := []string{"tmp"}
	removeDirectories(directories)

	fmt.Println("Have cleaned")
}
