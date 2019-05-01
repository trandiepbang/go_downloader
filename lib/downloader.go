package downloader

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// FileDownloader ...
type FileDownloader struct {
	Token  string
	Method string
}

// DataDownloader ...
type DataDownloader struct {
	Type string
	Data string
}

// NewDownloader ...
func NewDownloader(token string, method string) *FileDownloader {
	return &FileDownloader{
		Token:  token,
		Method: method,
	}
}

// WatchProgress ...
func (d *FileDownloader) WatchProgress(done chan int64, filePath string) {
	stop := false

	for {
		select {
		case <-done:
			stop = true
		default:
			file, err := os.Open(filePath)
			if err != nil {
				log.Fatal(err)
			}

			fi, err := file.Stat()
			if err != nil {
				log.Fatal(err)
			}

			size := fi.Size()

			if size == 0 {
				size = 1
			}

			data := &DataDownloader{
				Data: strconv.FormatInt(size, 10),
				Type: "message",
			}
			e, err := json.Marshal(data)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(e))
		}

		if stop {
			break
		}

		time.Sleep(time.Second)
	}
}

// DownloadFile ...
func (d *FileDownloader) DownloadFile(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Init req
	req, err := http.NewRequest(d.Method, url, nil)
	if err != nil {
		return fmt.Errorf("error init http new request")
	}

	// Set Data
	if d.Token != "" {
		req.Header.Add("Authorization", "Bearer "+d.Token)
	}

	// Get the data
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	done := make(chan int64)
	go d.WatchProgress(done, filepath)

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	m, err := io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	done <- m

	return nil
}
