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

	if argPathToClean == "/" {
		os.Exit(1)
	} else {
		for _, fileName := range pathDir {
			os.RemoveAll(path.Join([]string{(argPathToClean), fileName.Name()}...))
		}
	}
}

func groupUpFiles(argGroupUpFiles string) {
	extensions := []string{
		".txt",
		".ini",
		".md",
		".doc",
		".docx",
		".rtf",
		".pdf",
		".xlsx",
		".xls",
		".pptx",
		".zip",
		".7z",
		".gz",
		".bz",
		".gzip",
		".bzip",
		".iso",
		".mkv",
		".mov",
		".mp4",
		".bmp",
		".jpg",
		".png",
		".exe",
		".msi",
		".msu",
		".conf",
		".cfg",
		".net",
		".deny",
		".allow",
	}

	files, err := os.ReadDir(argGroupUpFiles)
	check(err)

	if argGroupUpFiles == "/" {
		os.Exit(1)
	} else {
		for _, file := range files {
			if strings.Contains(os.Args[0], file.Name()) {
				continue
			} else {
				for _, ext := range extensions {
					if strings.HasSuffix(file.Name(), ext) {
						newDir := argGroupUpFiles + "/" + "ALL" + strings.ToUpper(ext)
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
