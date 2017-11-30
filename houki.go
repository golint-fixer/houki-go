package houki

import (
	"fmt"
	"os"
	"sync"

	"github.com/apcera/termtables"
	"github.com/segmentio/go-prompt"
)

// Houki type for houki
type Houki struct{}

func (h *Houki) reCreateDirectory(directory string, wg *sync.WaitGroup) {
	defer wg.Done()
	if err := os.RemoveAll(directory); err != nil {
		fmt.Println(err)
	} else {
		os.Mkdir(directory, 0775)
	}
}

// RemoveDirectories remove specified directories
func (h *Houki) RemoveDirectories(directories []string) {
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
		go h.reCreateDirectory(directory, &wg)
	}
	wg.Wait()
	fmt.Println("Have cleaned")
}
