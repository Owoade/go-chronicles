package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/build", func(c *gin.Context) {
		// SSE headers
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")

		flusher, ok := c.Writer.(http.Flusher)
		if !ok {
			c.String(500, "Streaming unsupported")
			return
		}

		logChan := make(chan string, 100)
		done := make(chan bool)

		home, _ := os.UserHomeDir()

		cmd := exec.Command(
			"railpack",
			"build",
			filepath.Join(home, "infracon-apps/rust-ms"),
			"--name",
			"railpack-api-image-2",
		)

		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()

		if err := cmd.Start(); err != nil {
			c.String(500, err.Error())
			return
		}

		readPipe := func(pipe io.ReadCloser) {
			scanner := bufio.NewScanner(pipe)
			scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

			for scanner.Scan() {
				logChan <- scanner.Text()
			}
		}

		go readPipe(stdout)
		go readPipe(stderr)

		go func() {
			cmd.Wait()
			done <- true
		}()

		for {
			select {
			case line := <-logChan:
				fmt.Fprintf(c.Writer, "data: %s\n\n", line)
				flusher.Flush()

			case <-done:
				fmt.Fprintf(c.Writer, "data: BUILD FINISHED\n\n")
				flusher.Flush()
				return

			case <-c.Request.Context().Done():
				return
			}
		}
	})

	r.Run(":5000")
}
