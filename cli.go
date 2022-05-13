package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

func fatal_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func clean_the_dir(directory_path string) {
	dir, err := ioutil.ReadDir(directory_path)
	fatal_err(err)
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{directory_path, d.Name()}...))
	}
}

func main() {
	clean_the_dir("/test/test/")
}
