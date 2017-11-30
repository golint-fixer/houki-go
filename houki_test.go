package houki

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func TestReCreateDirectory(t *testing.T) {
	var houki Houki
	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	houki.reCreateDirectory(dir, &wg)
	_, err = os.Stat(tmpfn)
	expect := fmt.Sprintf("stat %s: no such file or directory", tmpfn)
	if err.Error() != expect {
		t.Errorf("Expect %s is not exists, but not %s", tmpfn, err)
	}

	if _, err = os.Stat(dir); err != nil {
		t.Errorf("Expect %s is exists, but not %s", dir, err)
	}
}
