package main

import (
	"flag"
	"fmt"
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

func group_up_files(new_dir_name string) {

	print("New dir name " + new_dir_name)
}

func main() {
	directory_path := flag.String("clean", ".", "Clean target directory")
	target_dir_name := flag.String("dir", ".", "Dir to group up the files")
	flag.Parse()

	switch os.Args[1] {
	case "--clean":
		fmt.Print(*directory_path)
		clean_the_dir(*directory_path)
	case "--dir":
		fmt.Println(*target_dir_name)
		group_up_files(*target_dir_name)
	default:
		os.Exit(1)
	}
}
