// main.go
package main

import (
	"github.com/toophy/pangu/thread"
	"testing"
)

// Gogame framework version.
func TestMain(t *testing.T) {
	thread.GetMaster().Wait_thread_over()
}
