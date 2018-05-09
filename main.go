package main

import (
	"fmt"
	"sync"
	"time"
	"github.com/amortaza/ksl/js"
)

var g_wait sync.WaitGroup
var g_exit = false

func main() {

	go checkJsFiles()
	go runServer()

	fmt.Println("running...")
	g_wait.Wait()
	fmt.Println("bye")
}

func checkJsFiles() {

	g_wait.Add(1)

	for !g_exit {

		js.CheckJsFiles()

		time.Sleep(5 * time.Second)
	}

	g_wait.Done()
}

func runServer() {

	time.Sleep(5 * time.Second)

	g_exit = true
}