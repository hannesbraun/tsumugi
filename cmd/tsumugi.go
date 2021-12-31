package main

import (
	"bufio"
	"fmt"
	"github.com/hannesbraun/tsumugi/panrec"
	"os"
	"strings"
	"time"
)

const VERSION = "0.9.0"

func main() {
	fmt.Println("Tsumugi", VERSION)
	datFiles := os.Args[1:]
	stdin := bufio.NewReader(os.Stdin)

	for _, path := range datFiles {
		if !strings.HasSuffix(path, ".dat") {
			continue
		}

		// TODO readonly flag
		updateDat(path, stdin, false)
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
