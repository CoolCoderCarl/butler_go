package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func fatal_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var target_dir_name string = "ALL"

func clean_the_dir(args_cln_name string) {

	dir, err := ioutil.ReadDir(args_cln_name)
	fatal_err(err)
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{args_cln_name, d.Name()}...))
	}
}

func group_up_files(args_dir_name string) {

	files_extension := []string{
		".txt",
		".doc",
		".docx",
		".pdf",
		".xlsx",
		".bmp",
		".jpg",
		".rtf",
		".pptx",
		".conf",
		".cfg",
		".net",
		".deny",
		".allow",
	}

	files, err := ioutil.ReadDir(args_dir_name)
	fatal_err(err)

	for _, file := range files {
		for _, ext := range files_extension {
			if strings.HasSuffix(file.Name(), ext) {
				new_dir_path := args_dir_name + target_dir_name + strings.ToUpper(ext)
				// Skip if cli.* has been met
				if _, err := os.Stat(new_dir_path); os.IsNotExist(err) {
					err := os.Mkdir(new_dir_path, 0644)
					fatal_err(err)
				}
				err := os.Rename(args_dir_name+file.Name(), new_dir_path+"/"+file.Name())
				fatal_err(err)
			}
		}
	}
}

func main() {
	args_cln_name := flag.String("clean", ".", "Clean target directory")
	args_dir_name := flag.String("dir", ".", "Dir to group up the files")
	flag.Parse()

	switch os.Args[1] {
	case "--clean":
		fmt.Print(*args_cln_name)
		clean_the_dir(*args_cln_name)
	case "--dir":
		fmt.Println(*args_dir_name)
		group_up_files(*args_dir_name)
	default:
		os.Exit(1)
		// Add msg about err
	}
}
