package main

import (
	"log"
	"os"
	"path"
)

func cleanDir() {

	pathDir, err := os.ReadDir("/test/test/")
	if err != nil {
		log.Fatal(err)
	}

	for _, fileName := range pathDir {
		os.RemoveAll(path.Join([]string{("/test/test/"), fileName.Name()}...))
	}

}

func main() {
	cleanDir()
	// FUNC TO GROUP UP
}
