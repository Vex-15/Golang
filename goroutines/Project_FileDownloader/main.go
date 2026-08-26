package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func downloadFile(url, dest string) error {
	fileName := filepath.Base(url)
	filePath := filepath.Join(dest, fileName)

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	fmt.Printf("Downloading the file :D %s\n", url)
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		_ = os.Remove(filePath)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		_ = os.Remove(filePath)
		return fmt.Errorf("Bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Downloading %s took %s\n", fileName, time.Since(start))
	return nil

}

func sequentialDownloader(urls []string, destDir string) error {
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	start := time.Now()
	for _, url := range urls {
		if err := downloadFile(url, destDir); err != nil {
			fmt.Println("Error downloading", url, err)
			continue
		}
	}
	fmt.Printf("Downloading %d files took %s\n", len(urls), time.Since(start))
	return nil

}

type Result struct {
	URL      string
	FileName string
	Size     int64
	Duration time.Duration
	Error    error
}

func concurrentDownloader(urls []string, destDir string, maxConcurrent int) error {

	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}
	start := time.Now()
	results := make(chan Result)

	var wg sync.WaitGroup

	limiter := make(chan struct{}, maxConcurrent)
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			limiter <- struct{}{}
			defer func() {
				<-limiter
			}()

			start := time.Now()

			filename := filepath.Base(url)
			filePath := filepath.Join(destDir, filename)

			out, err := os.Create(filePath)
			if err != nil {
				results <- Result{URL: url, Error: err}
				return
			}
			defer out.Close()

			resp, err := http.Get(url)
			if err != nil {
				results <- Result{URL: url, Error: err}
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				results <- Result{
					URL:   url,
					Error: fmt.Errorf("bad status: %s", resp.Status),
				}
				return
			}

			size, err := io.Copy(out, resp.Body)
			if err != nil {
				results <- Result{URL: url, Error: err}
				return
			}

			timeSince := time.Since(start)
			results <- Result{URL: url, FileName: filename, Size: size, Duration: timeSince}
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var totalSize int64
	var errors []error
	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error downloading %s: %s\n", result.URL, result.Error.Error())
			errors = append(errors, result.Error)
		} else {
			totalSize += result.Size
			fmt.Printf("Downloaded %s (%d bytes) in %s\n", result.FileName, result.Size, result.Duration)
		}
	}
	startedSince := time.Since(start)
	fmt.Printf("All downloads completed in %s, total size: %d bytes\n", startedSince, totalSize)
	if len(errors) > 0 {
		return fmt.Errorf("%d downloads failed", len(errors))
	}
	return nil

}
func main() {

	// url := "https://chameleonmemes.com/wp-content/uploads/2023/07/I-survived.jpg"

	// err := downloadFile(url, "./")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// Above one is input for a single file download

	urls := []string{
		"https://tse3.mm.bing.net/th/id/OIP.I76wCCHMWyK9EqVb-5GP_QHaHP?rs=1&pid=ImgDetMain&o=7&rm=3",
		"https://content.imageresizer.com/images/memes/Cool-cat-meme-2.jpg",
	}

	err := concurrentDownloader(urls, "./", 3)
	if err != nil {
		println(err)
		return
	}

	log.Println("Done")
}
