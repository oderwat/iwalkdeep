// ideepwalk project ideepwalk.go
package iwalkdeep

import (
	//"errors"
	"io"
	"os"
	"path/filepath"
)

type WalkDeepFunc func(dir string)

func walkdeep(path string, fun WalkDeepFunc) string {
	d, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer d.Close()

	// I am not sorting and just follow directories here!
	// This is because when it works on filesystems
	// with many files and directories it would need
	// to store all that data while diving into the
	// subfolders recursively!

	for {
		x, err := d.Readdir(1)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		f := x[0]

		if f.IsDir() {
			fun(walkdeep(filepath.Join(path, f.Name()), fun))
		}
	}
	return path
}

func WalkDeep(root string, fun WalkDeepFunc) {
	fun(walkdeep(root, fun))
}

// This is the "Iterator" implemented with a go chanel
func IWalkDeep(root string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		WalkDeep(root, func(dir string) {
			ch <- dir
		})
	}()
	return ch
}
