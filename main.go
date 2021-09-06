package main

import (
	"bytes"
	"encoding/json"
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
	var buf bytes.Buffer
	// JSONフォーマットでないパターンもたくさんあるのでエラーは捨てる
	json.Indent(&buf, []byte(text), "", "  ")
	clipboard.WriteAll(buf.String())
}
