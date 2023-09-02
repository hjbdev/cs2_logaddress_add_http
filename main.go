package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"bytes"

	"github.com/leaanthony/clir"
)

func command(fileName string, url string) {
	// monitor file for changes, retrieve new lines, POST to url as raw body
	var buffer bytes.Buffer

	interval := time.Second / 128 // 128 tick
	ticker := time.NewTicker(interval)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// get file size
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Error getting file info: %v\n", err)
		return
	}

	previousSize := fileInfo.Size()

	defer ticker.Stop()

	for range ticker.C {
		// read new lines from file
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Printf("Error getting file info: %v\n", err)
			return
		}
		newSize := fileInfo.Size()
		if newSize > previousSize {
			// seek to previous size
			file.Seek(previousSize, 0)

			fileBytes := make([]byte, newSize-previousSize)

			_, err := file.Read(fileBytes)

			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
				return
			}

			// seek to end of file
			file.Seek(newSize, 0)

			// send new lines to url
			fmt.Printf("Sending new lines to %s\n", url)

			// turn bytes into a buffer
			buffer.Write(fileBytes)

			fmt.Println(buffer.String())

			go http.Post(url, "text/plain", bytes.NewBuffer(buffer.Bytes()))

			buffer.Reset()

			previousSize = newSize
		}
	}

	// go http.Post(url, "text/plain", &buffer)
}

func main() {
	cli := clir.NewCli("cs2_logaddress_add_http", "A workaround for the current lack of logaddress_add_http in Counter-Strike 2", "1.0.0")

	fileName := "server.log"
	cli.StringFlag("file", "Server log file location", &fileName)

	url := ""
	cli.StringFlag("url", "URL to send POST requests to", &url)

	cli.Action(func() error {
		if url == "" {
			fmt.Printf("Error: url is required\n")
			return nil
		}

		fmt.Printf("Monitoring %s for changes, sending new lines to %s\n", fileName, url)

		command(fileName, url)

		return nil
	})

	if err := cli.Run(); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
	}
}
