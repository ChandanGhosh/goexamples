// Here we shall be looking at the errorgroup package(https://pkg.go.dev/golang.org/x/sync/errgroup).
// this package helps in grouping and managing goroutine errors effectively. To demonstrate,
// we have taken a simple implementation of counting words in text files
// using linux binary "wc". If you genrally execute "wc --words atomic.html" from terminal
// you shall see a output like this "4810 atomic.html". Here the first in the
// string is count of words. If there is any error in opening a file, file does not exists,
// buffer parsing issues etc. the wordCount will return error and the program will exit
// with the exact error. This way of handling error in concurrent gorouting execution always ensures
// program terminates in case of any errors generated. The "sync/atomic" package usually communicate
// by sharing memory like here(&totalWordCount). This is a low-level package which should be used
// with special care. Usually when in highly concurrent scenerio, it is better to
// share memory by communicating(using channel) instead of communicating by sharing memory address (&totalwordCount).

package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

func wordCount(filepath string) (int, error) {
	out, err := exec.Command("wc", "--words", filepath).Output()
	if err != nil {
		return 0, fmt.Errorf("error getting words count in the file %s is : %s", filepath, err)
	}
	wc, err := strconv.Atoi(strings.Fields(string(out))[0])
	if err != nil {
		return 0, fmt.Errorf("error parsing the output for file %s is %s", filepath, err)
	}

	return wc, nil
}

func main() {
	var totalWordCount int64
	var errGrp errgroup.Group
	files := []string{
		"atomic.html",
		"errgroup.html",
	}
	for _, file := range files {
		file := file
		errGrp.Go(func() error {
			wc, err := wordCount(file)
			if err != nil {
				return err
			}
			atomic.AddInt64(&totalWordCount, int64(wc))
			return nil
		})
	}

	if err := errGrp.Wait(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Total word count of all the files : %v", totalWordCount)
}
