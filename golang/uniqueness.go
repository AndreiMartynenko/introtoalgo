package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func uniq(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(os.Stdin)
	var prev string
	// alreadySeen := make(map[string]bool)
	for in.Scan() {
		txt := in.Text()
		if txt == prev {
			continue
		}
		if txt < prev {
			// panic("file not sorted")
			return fmt.Errorf("file not sorted")
		}
		// if _, found := alreadySeen[txt]; found {
		// 	continue
		// }
		//remember that we've already seen
		// alreadySeen[txt] = true
		prev = txt
		fmt.Fprintln(output, txt)
	}
	return nil

}

func main() {
	err := uniq(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}

}
