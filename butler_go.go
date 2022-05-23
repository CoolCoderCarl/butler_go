package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
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

func cleanTheDir(argPathToClean string) {

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

func createArchive() {

	files, err := os.ReadDir("/test/test/")
	check(err)

	for _, file := range files {
		fmt.Println(file.Name())
		archive, err := os.Create("archive.zip")
		check(err)
		defer archive.Close()
		zipWriter := zip.NewWriter(archive)
		file, err := os.Open(file.Name())
		check(err)
		defer file.Close()
		w_file, err := zipWriter.Create(file.Name())
		check(err)
		if _, err := io.Copy(w_file, file); err != nil {
			panic(err)
		}
		zipWriter.Close()
	}
}

func main() {
	argPathToClean := flag.String("clean", ".", "Remove all files from target directory")
	argGroupUpFiles := flag.String("group", ".", "Group up all files in target directory")
	// argCreateArchive := flag.String("archive", ".", "Create archive from target directory")
	flag.Parse()

	switch os.Args[1] {
	case "--clean":
		fmt.Println(*argPathToClean)
		cleanTheDir(*argPathToClean)
	case "--group":
		fmt.Println(*argGroupUpFiles)
		groupUpFiles(*argGroupUpFiles)
		// case "--archive":
		// 	fmt.Println(*argCreateArchive)
		// 	createArchive(*argCreateArchive)
	}
	// createArchive()
}
