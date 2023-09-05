package main

import (
	"bufio"
	"container/heap"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

// ReadLines reads lines from an io.Reader and returns them as a slice of strings.
func ReadLines(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// SortLines sorts a slice of strings in ascending order.
func SortLines(lines []string) {
	sort.Strings(lines)
}

// WriteLines writes a slice of strings to an io.Writer.
func WriteLines(w io.Writer, lines []string) error {
	writer := bufio.NewWriter(w)

	for _, line := range lines {
		_, err := fmt.Fprintln(writer, line)
		if err != nil {
			return err
		}
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}

// WriteLinesToTempFile writes lines to a temporary file and returns its name.
func WriteLinesToTempFile(lines []string) (string, error) {
	tmpFile, err := os.CreateTemp("", "sort-chunk")
	if err != nil {
		return "", err
	}
	defer tmpFile.Close()

	err = WriteLines(tmpFile, lines)
	if err != nil {
		return "", err
	}

	return tmpFile.Name(), nil
}

// LineRecord stores a line and its origin file.
type LineRecord struct {
	line string
	file *os.File
}

// LineHeap is a priority queue for LineRecord based on line string.
type LineHeap []LineRecord

func (h LineHeap) Len() int           { return len(h) }
func (h LineHeap) Less(i, j int) bool { return h[i].line < h[j].line }
func (h LineHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *LineHeap) Push(x interface{}) {
	*h = append(*h, x.(LineRecord))
}

func (h *LineHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MergeTempFiles merges multiple sorted temp files into a single sorted output file.
func MergeTempFiles(tempFiles []string, outFile *os.File) error {
	h := &LineHeap{}
	heap.Init(h)
	fileReaders := make([]*bufio.Reader, len(tempFiles))

	// Open all temp files and read the first line from each, add to heap.
	for i, fileName := range tempFiles {
		file, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer file.Close()
		fileReaders[i] = bufio.NewReader(file)

		line, err := fileReaders[i].ReadString('\n')
		if err != nil {
			return err
		}
		heap.Push(h, LineRecord{line: line, file: file})
	}

	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	for h.Len() > 0 {
		smallest := heap.Pop(h).(LineRecord)
		_, err := writer.WriteString(smallest.line)
		if err != nil {
			return err
		}

		newLine, err := bufio.NewReader(smallest.file).ReadString('\n')
		if err == nil {
			heap.Push(h, LineRecord{line: newLine, file: smallest.file})
		} else if err != io.EOF {
			return err
		}
	}

	return nil
}

func main() {
	// Flag for external sorting
	useExternalSort := flag.Bool("external", false, "Use external sorting")
	flag.Parse()

	file, err := os.Open("../input4.txt")
	if err != nil {
		log.Printf("Error reading file: %s", err)
		return
	}
	defer file.Close()

	lines, err := ReadLines(file)
	if err != nil {
		log.Printf("Error reading lines: %s", err)
		return
	}

	if *useExternalSort {
		// 1. Sort lines and write them into multiple temp files
		chunkSize := 1 // Or whatever size you want each chunk to be
		var tempFiles []string
		for i := 0; i < len(lines); i += chunkSize {
			end := i + chunkSize
			if end > len(lines) {
				end = len(lines)
			}
			chunk := lines[i:end]
			SortLines(chunk) // Sort each chunk

			tempFileName, err := WriteLinesToTempFile(chunk)
			if err != nil {
				log.Printf("Error writing to temp file: %s", err)
				return
			}

			tempFiles = append(tempFiles, tempFileName)
		}

		// 2. Merge these temp files into a single sorted file
		outFile, err := os.Create("../sorted-file/sorted.txt")
		if err != nil {
			log.Printf("Error creating sorted file: %s", err)
			return
		}
		defer outFile.Close()

		err = MergeTempFiles(tempFiles, outFile)
		if err != nil {
			log.Printf("Error merging temp files: %s", err)
			return
		}

	} else {
		SortLines(lines)

		outFile, err := os.Create("../sorted-file/sorted.txt")
		if err != nil {
			log.Printf("Error writing to file: %s", err)
			return
		}
		defer outFile.Close()

		err = WriteLines(outFile, lines)
		if err != nil {
			log.Printf("Error writing lines: %s", err)
		}
	}
}
