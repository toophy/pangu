// main.go
package main

import (
	"fmt"
	"github.com/toophy/pangu/help"
	"github.com/toophy/pangu/thread"
	"os"
	"runtime/pprof"
)

// Gogame framework version.
const (
	VERSION = "0.0.2"
)

func main() {
	// 检查log目录
	if !help.IsExist(thread.LogDir) {
		os.MkdirAll(thread.LogDir, os.ModeDir)
	}

	// 创建pprof文件
	f, err := os.Create(thread.LogDir + "/" + thread.ProfFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// 启动并等待主线程
	thread.GetWorld().Wait_thread_over()
}
