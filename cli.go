package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func wordcounter(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	scanner.Split(bufio.ScanWords)

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}

func main() {
	fmt.Print("Input words, then CTRL+C:")
	fmt.Println("Total words:", wordcounter(os.Stdin))
}
