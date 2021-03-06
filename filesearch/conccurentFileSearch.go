package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	count     int
	matches   []string
	waitgroup = sync.WaitGroup{}
	lock      = sync.Mutex{}
)

func fileSearch(root string, filename string) {
	fmt.Println("Searching in ", root)
	files, _ := ioutil.ReadDir(root)
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))
			lock.Unlock()
		}
		if file.IsDir() {
			waitgroup.Add(1)
			go fileSearch(filepath.Join(root, file.Name()), filename)
		}
	}
	waitgroup.Done()
}

func main() {
	start := time.Now()
	waitgroup.Add(1)
	go fileSearch("/", ".pdf")
	waitgroup.Wait()
	for _, file := range matches {
		fmt.Println("Matched", file)
		count++
	}

	duration := time.Since(start)

	switch count {
	case 1:
		fmt.Println("Found 1 result.")
	case 0:
		fmt.Println("No results found")
	default:
		fmt.Printf("Found %d results.\n", count)
	}

	fmt.Printf("Time required for search: %f seconds.\n", duration.Seconds())

}
