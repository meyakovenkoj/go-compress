package typing

import (
	"flag"
	"fmt"

	"github.com/gabriel-vasile/mimetype"
)

const (
	Unknown int = iota
	Executable
	Binary
	Image
	Text
	Json
)

type FileStats [6]float64

func Normalize(stat FileStats) FileStats {
	sum := 0.0
	sum += stat[Executable] + stat[Binary] + stat[Image] + stat[Text] + stat[Json] + stat[Unknown]
	newStats := FileStats{}
	newStats[Executable] = stat[Executable] / sum
	newStats[Binary] = stat[Binary] / sum
	newStats[Image] = stat[Image] / sum
	newStats[Text] = stat[Text] / sum
	newStats[Json] = stat[Json] / sum
	newStats[Unknown] = stat[Unknown] / sum
	return newStats
}

func SetStat(typeOf int, size int64, stat *FileStats) {
	switch typeOf {
	case Unknown:
		stat[Unknown] += float64(size)
		return
	case Binary:
		stat[Binary] += float64(size)
		return
	case Image:
		stat[Image] += float64(size)
		return
	case Executable:
		stat[Executable] += float64(size)
		return
	case Json:
		stat[Json] += float64(size)
		return
	case Text:
		stat[Text] += float64(size)
		return
	}
}

func Match(filename string) int {
	mimetype.SetLimit(1024 * 1024)
	mtype, err := mimetype.DetectFile(filename)
	if err != nil {
		fmt.Println(err)
		return Unknown
	}
	if mtype.Is("application/octet-stream") {
		return Binary
	} else if mtype.Is("image/tiff") {
		return Image
	} else if mtype.Is("application/x-mach-binary") || mtype.Is("application/vnd.microsoft.portable-executable") {
		return Executable
	} else if mtype.Is("application/json") {
		return Json
	} else if mtype.Is("text/plain") || mtype.Parent().Is("text/plain") {
		return Text
	} else {
		return Unknown
	}
}

func main() {
	filePtr := flag.String("file", "", "file to check")
	flag.Parse()

	if *filePtr == "" {
		fmt.Println("need to set file name with flag -file")
		return
	}
	Match(*filePtr)
}
