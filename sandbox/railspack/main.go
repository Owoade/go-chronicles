package main

import (
	"bufio"
	"fmt"
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

		flusher, ok := c.Writer.(gin.ResponseWriter)
		if !ok {
			return
		}

		logChan := make(chan string)
		done := make(chan bool)

		home, _ := os.UserHomeDir()

		cmd := exec.Command(
			"railpack",
			"build",
			filepath.Join(home, "infracon-apps/rust-ms"),
			"--name",
			"railpack-api-image",
		)

		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()

		_ = cmd.Start()

		// read stdout
		go func() {
			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				logChan <- scanner.Text()
			}
		}()

		// read stderr
		go func() {
			scanner := bufio.NewScanner(stderr)
			for scanner.Scan() {
				logChan <- scanner.Text()
			}
		}()

		// wait for command
		go func() {
			cmd.Wait()
			done <- true
		}()

		// SSE loop
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

	r.Run(":3000")
}
