package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gopl.io/ch8/thumbnail"
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

func makeThumbnails6(filenames <-chan string, sizes chan<- int64) {
	var wg sync.WaitGroup

	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}
	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()
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
	// makeThumbnails4(paths)
	// makeThumbnails5(paths)

	// for makeThumbnails6
	ch := make(chan string, len(paths))
	sizes := make(chan int64, len(paths))

	go makeThumbnails6(ch, sizes)
	for _, path := range paths {
		fmt.Println(path)
		ch <- path
	}
	close(ch)

	var total int64
	// sizes channel forces the main channel to wait
	for size := range sizes {
		total += size
	}

	fmt.Printf("Done, bytes written: %d\n", total)

}
