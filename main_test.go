package main

import (
	"bytes"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRealMain(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	go func() {
		realMain(&buf)
	}()

	// Wait for the goroutine to start and run the server
	time.Sleep(10 * time.Millisecond)

	assert.Equal(t, 1, realMain(&buf))
	assert.Contains(t, buf.String(), "GIN HTTP server running on port")
}
