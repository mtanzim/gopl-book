package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mtanzim/gopl-book/ch8/thumbnail/thumbnail"
)

// make thumbnails serially
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}

}

// Incorrect, returns before creating thumbnails
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f)
	}
}

// make thumbnails in parallel with an unbuffered channel
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		// note the usage of function closure here
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}
	for range filenames {
		// wait for all goroutines to complete
		<-ch
	}
}

// pass back errors incorrectly
// demonstrates common goroutine leaks
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		// note the usage of function closure here
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}
	for range filenames {
		if err := <-errors; err != nil {
			// Note the goroutine leak here
			// more error messages may be produced with no receiver to drain them
			return err
		}
	}
	return nil
}

// fixes the goroutine leak from makeThumbnails4
// with buffered channels
func makeThumbnails5(filenames []string) ([]string, error) {

	type item struct {
		thumbfile string
		err       error
	}
	// make a buffered channel with max capacity
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		// note the usage of function closure here
		go func(f string) {
			thumbfile, err := thumbnail.ImageFile(f)
			ch <- item{thumbfile: thumbfile, err: err}
		}(f)
	}

	thumbfiles := []string{}
	for range filenames {
		it := <-ch
		if it.err != nil {
			// even with the early return here
			// since the channel is not full, there is no goroutine leak
			// in other words, no goroutine will be blocked from sending or receiving since the buffer is not full
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

func getImagePaths() []string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	filenames := []string{"a", "b", "c", "d"}
	imagePaths := []string{}
	for _, name := range filenames {
		imagePaths = append(imagePaths, fmt.Sprintf("%s/ch8/thumbnail/assets/%s.jpg", path, name))
	}
	return imagePaths
}

func main() {

	paths := getImagePaths()
	// makeThumbnails2(paths)
	// makeThumbnails3(paths)
	makeThumbnails4(paths)
}
