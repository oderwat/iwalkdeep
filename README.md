PACKAGE DOCUMENTATION

package iwalkdeep
    import "github.com/oderwat/iwalkdeep"

    The package "ideepwalk" helps to iterates over all directories in a give
    path in bottom to top order. This is helpfull in many use cases where
    you for example want to clean up empty directories after
    processing/moving files in the hierarchie.

    As speciality the packaged exposes this functionality also as an easy to
    use iterator like in the following example

	for dir := range IWalkDeep("/homes") {
	  fmt.print(dir)
	}


FUNCTIONS

func IWalkDeep(root string) <-chan string
    This is the "Iterator" which can eaasily be used for dir := range
    IWalkDeep(path) {}

func WalkDeep(root string, fun WalkDeepFunc)
    Walk all directories in root from bottom to top calling WalkDeepFunc for
    every encountered directory


TYPES

type WalkDeepFunc func(dir string)
    This function is called by WalkDeep for each directory on its way up to
    the root of the given path




