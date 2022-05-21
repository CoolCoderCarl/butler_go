package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func cleanDir(argPathToClean string) {

	pathDir, err := os.ReadDir(argPathToClean)
	check(err)

	for _, fileName := range pathDir {
		os.RemoveAll(path.Join([]string{(argPathToClean), fileName.Name()}...))
	}

}

func groupUpFiles(argGroupUpFiles string) {
	extensions := []string{".docx", ".pptx", ".txt", ".xlsx"}

	files, err := os.ReadDir(argGroupUpFiles)
	check(err)

	for _, file := range files {
		for _, ext := range extensions {
			if strings.HasSuffix(file.Name(), ext) {
				newDir := argGroupUpFiles + "/" + "ALL_" + strings.ToUpper(ext)
				if _, err := os.Stat(newDir); os.IsNotExist(err) {
					err := os.Mkdir(newDir, 0644)
					check(err)
				}
				err := os.Rename(argGroupUpFiles+"/"+file.Name(), newDir+"/"+file.Name())
				check(err)
			}
		}
	}
}

func main() {
	argPathToClean := flag.String("clean", ".", "Remove all files from target directory")
	argGroupUpFiles := flag.String("group", ".", "Group up all files in target directory")
	flag.Parse()

	switch os.Args[1] {
	case "--clean":
		fmt.Println(*argPathToClean)
		cleanDir(*argPathToClean)
	case "--group":
		fmt.Println(*argGroupUpFiles)
		groupUpFiles(*argGroupUpFiles)
	}
}
