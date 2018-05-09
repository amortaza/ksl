package js

import (
	"time"
	"os"
	"fmt"
	"io/ioutil"
)

type JsFile struct {

	filename    string
	updatedTime time.Time
	js          string
}

var g_jsFilesByFilename = make(map[string] *JsFile)

func CheckJsFiles() {

	for _, jsfile := range g_jsFilesByFilename {

		jsfile.check()
	}
}

func LoadJsFile(filename string) *JsFile {

	jsfile, ok := g_jsFilesByFilename[filename]

	if !ok {

		jsfile = &JsFile{filename: filename}

		jsfile.check()

		g_jsFilesByFilename[filename] = jsfile
	}

	return jsfile
}

func (jsfile *JsFile) check() bool {

	info, err := os.Stat(jsfile.filename)

	if err != nil {
		fmt.Println(err, " Unable to open JS file ", jsfile.filename)
		panic("Unable to open JS file " + jsfile.filename)
	}

	if info == nil {
		fmt.Println("Unable to open file ", jsfile.filename)
		panic("Unable to open JS file " + jsfile.filename)
	}

	updatedTime :=  info.ModTime()

	if jsfile.updatedTime != updatedTime {

		jsfile.updatedTime = updatedTime

		buf, _ := ioutil.ReadFile(jsfile.filename)

		jsfile.js = string(buf)

		fmt.Println("loaded js file..." + jsfile.filename)

		return true
	}

	return false
}


