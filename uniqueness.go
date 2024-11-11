package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	var prev string
	// alreadySeen := make(map[string]bool)
	for in.Scan() {
		txt := in.Text()
		if txt == prev {
			continue
		}
		if txt < prev {
			panic("file not sorted")
		}
		// if _, found := alreadySeen[txt]; found {
		// 	continue
		// }
		//remember that we've already seen
		// alreadySeen[txt] = true
		prev = txt
		fmt.Println(txt)
	}

}
