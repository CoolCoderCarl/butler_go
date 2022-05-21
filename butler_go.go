package main

import (
	"flag"
	"log"
	"os"
	"path"
)

func cleanDir(argPathToClean string) {

	pathDir, err := os.ReadDir(argPathToClean)
	if err != nil {
		log.Fatal(err)
	}

	for _, fileName := range pathDir {
		os.RemoveAll(path.Join([]string{(argPathToClean), fileName.Name()}...))
	}

}

func main() {
	argPathToClean := flag.String("clean", ".", "Remove all files from target directory")
	flag.Parse()

	switch os.Args[1] {
	case "--clean":
		print(*argPathToClean)
		cleanDir(*argPathToClean)
	}

	// FUNC TO GROUP UP
}
