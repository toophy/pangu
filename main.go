// main.go
package main

import (
	"github.com/toophy/Gogame/thread"
)

// Gogame framework version.
const VERSION = "0.0.2"

func main() {
	thread.GetMaster().Wait_thread_over()
}
