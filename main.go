package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sink_file_count int = 0

type SinkFile struct {
	id         int
	filename   string
	url        string
	downloaded string
}

func (this SinkFile) Print() {
	fmt.Printf("[%d] %s (%d): %s (%d)\n", this.id, this.filename, this.LocalFileLength(), this.url, this.DownloadedLength())
}

func (this SinkFile) LocalFileLength() int64 {
	file, err := os.Stat(this.filename)
	if err != nil {
		return 0
	}
	return file.Size()
}

func (this SinkFile) DownloadedLength() int {
	return len(this.downloaded)
}

func (this SinkFile) Download() {
	this.downloaded = "Downloaded text"
}

func (this SinkFile) Update() {
	//
}

func SinkFileFactory(line string) SinkFile {
	args := strings.Split(line, "=")
	var sf SinkFile
	sf.id = sink_file_count
	sf.filename = strings.TrimSpace(args[0])
	sf.url = strings.TrimSpace(args[1])
	sink_file_count++
	return sf
}

func main() {
	filename := "sink.sink"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("File not found: %s\n", filename)
		return
	}

	sinks := []SinkFile{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			sf := SinkFileFactory(line)
			sinks = append(sinks, sf)
		}
	}

	for _, sink := range sinks {
		sink.Print()
	}

	if err := scanner.Err(); err != nil {
		return
	}
	return
}
