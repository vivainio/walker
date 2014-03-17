package walker

import "os"
import "fmt"

// should return true if you want to descend further, false if stop
type FastWalkCallback func(string, []os.FileInfo) bool

func WalkOne(pth string, cb FastWalkCallback) {
	dir, err := os.Open(pth)
	if err != nil {
		pe := err.(*os.PathError)
		fmt.Printf("Path error: %s (%s)\n", pe, pth)
		return
	}

	fis, err := dir.Readdir(-1)

	//fmt.Printf("%s", pth)
	r := cb(pth, fis)
	if !r {
		return
	}

	for _, fi := range fis {
		if fi.IsDir() {
			WalkOne(pth+"/"+fi.Name(), cb)
		}
		//fmt.Println(fi.Name())
	}

}
