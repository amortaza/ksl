package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/amortaza/ksl/js"
)

func SysJsHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	var filename = vars["js_filename"]

	filename = "github.com/amortaza/ksl/scripts/sys_js/" + filename + ".js"

	jsfile := js.LoadJsFile(filename)

	queries := r.URL.Query()

	params := make(map[string]string)

	for k, v := range queries {
		params[k] = v[0]
	}

	fmt.Println("running js " + jsfile.Filename)

	js.RunJs(jsfile, params)

	w.WriteHeader(http.StatusOK)
}

func QuitHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("quit request received")
	g_exit = true
	w.WriteHeader(http.StatusOK)
}
