package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/hannesbraun/tsumugi/panrec"
	"os"
	"strings"
	"time"
)

const VERSION = "0.9.1"

func main() {
	var readOnly bool
	flag.BoolVar(&readOnly, "readonly", false, "don't ask for a new title")
	flag.Parse()
	flag.Args()

	fmt.Println("Tsumugi", VERSION)
	datFiles := os.Args[1:]
	stdin := bufio.NewReader(os.Stdin)

	for _, path := range datFiles {
		if !strings.HasSuffix(path, ".dat") {
			continue
		}

		updateDat(path, stdin, readOnly)
	}
}

func updateDat(path string, stdin *bufio.Reader, readOnly bool) {
	hLine()
	fmt.Println("File:", path)

	data := panrec.Read(path)

	fmt.Println("Timestamp:", time.Unix(int64(data.Timestamp), 0).String())
	fmt.Println("Title:", data.Title)
	fmt.Println("TV Channel:", data.Channel)
	fmt.Println("Language:", data.Language)
	fmt.Println("Viewed:", data.Viewed)
	fmt.Println()

	if readOnly {
		return
	}

	fmt.Print("New title: ")
	title, err := stdin.ReadString('\n')
	if err != nil {
		panic(err)
	}
	title = strings.TrimSpace(title)
	if len(title) > panrec.TitleLength {
		title = title[:panrec.TitleLength-1] // Minus one to have at least one null byte
	}
	data.Title = title

	if len(data.Title) > 0 {
		panrec.UpdateTitle(path, data)
	}
}

func hLine() {
	fmt.Println("--------------------------------------------------------------------------------")
}
