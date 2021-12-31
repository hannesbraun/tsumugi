package panrec

import (
	"encoding/binary"
	"os"
)

const (
	TitleLength = 124
	TitleOffset = 0x24
)

type Metadata struct {
	Timestamp uint32
	Title     string
	Channel   string
	Language  string
	Viewed    bool
}

func Read(path string) Metadata {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	timestamp := binary.BigEndian.Uint32(file[0x08:0x0c])
	title := string(file[TitleOffset : TitleOffset+TitleLength])
	channel := string(file[0x12c:0x150])
	language := string(file[0x170:0x173])
	viewed := file[0x173] == 0

	return Metadata{
		timestamp,
		title,
		channel,
		language,
		viewed,
	}
}

func UpdateTitle(path string, data Metadata) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	title := []byte(data.Title)

	// Delete title
	for i := 0; i < TitleLength; i++ {
		if i < len(title) {
			file[TitleOffset+i] = title[i]
		} else {
			file[TitleOffset+i] = 0
		}
	}

	err = os.WriteFile(path, file, 0644)
	if err != nil {
		panic(err)
	}
}
