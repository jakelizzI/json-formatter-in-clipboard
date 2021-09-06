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
	err2 := json.Indent(&buf, []byte(text), "", "  ")
	if err2 != nil {
		// JSONフォーマットでないパターンもたくさんあるのでエラーは捨てる
		return
	}
	clipboard.WriteAll(buf.String())
}
