package main

import (
	"fmt"
	"sync"
	"time"
	"github.com/amortaza/ksl/js"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/robertkrimen/otto"
)

var g_wait sync.WaitGroup
var g_exit = false

func main() {

	vm := otto.New()
	foo, _ := vm.Compile("", "function foo(){console.log('this is foo');}");


	vm.Run(foo);

	vm.Run("foo()")

	v1 := vm.Copy();

	boo, _ := vm.Compile("", "function foo(){console.log('this is boo');}");
	vm.Run(boo);

	v2 := vm.Copy();

	v1.Run("foo(); boo();")
	v2.Run("foo(); boo();")

	//go checkJsFiles()
	//go runServer()
	//
	//fmt.Println("waiting on threads")
	//g_wait.Wait()
	//fmt.Println("bye")
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

	router := mux.NewRouter()
	router.HandleFunc("/quit", QuitHandler)
	router.HandleFunc("/sys_js/{js_filename}", SysJsHandler)

	http.Handle("/", router)

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 3600 * 24 * time.Second,
		ReadTimeout:  3600 * 24 * time.Second,
	}

	fmt.Println("listening...")
	server.ListenAndServe()
}

