// Package iwalkdeep is an iterator for the filesystem
//
// It helps to iterates over all directories in a give path in bottom to top order.
// This is helpfull in many use cases where you for example want to clean up empty directories
// after processing/moving files in the hierarchie.
//
// As speciality the packaged exposes this functionality also as an
// easy to use iterator like in the following example:
//
//  for dir := range IWalkDeep("/homes") {
//    fmt.print(dir)
//  }
package iwalkdeep

import (
	//"errors"
	"io"
	"os"
	"path/filepath"
)

// WalkDeepFunc is called by WalkDeep for each directory
// on its way up to the root of the given path
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

// WalkDeep walkes all directories in root from bottom to top calling
// WalkDeepFunc for every encountered directory
func WalkDeep(root string, fun WalkDeepFunc) {
	fun(walkdeep(root, fun))
}

// IWalkDeep is the "Iterator" which can eaasily be used
// for dir := range IWalkDeep(path) {}
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
