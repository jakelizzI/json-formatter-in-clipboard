package main

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	text, err := clipboard.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(text)
}
