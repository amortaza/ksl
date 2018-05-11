package js

import (
	"time"
	"os"
	"fmt"
	"io/ioutil"
)

type JsFile struct {

	Filename    string
	updatedTime time.Time
	JsScript    string
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

		jsfile = &JsFile{Filename: filename}

		jsfile.check()

		g_jsFilesByFilename[filename] = jsfile
	}

	return jsfile
}

func (jsfile *JsFile) check() bool {

	info, err := os.Stat(jsfile.Filename)

	if err != nil {
		fmt.Println(err, " Unable to open JS file ", jsfile.Filename)
		panic("Unable to open JS file " + jsfile.Filename)
	}

	if info == nil {
		fmt.Println("Unable to open file ", jsfile.Filename)
		panic("Unable to open JS file " + jsfile.Filename)
	}

	updatedTime :=  info.ModTime()

	if jsfile.updatedTime != updatedTime {

		jsfile.updatedTime = updatedTime

		buf, _ := ioutil.ReadFile(jsfile.Filename)

		jsfile.JsScript = string(buf)

		fmt.Println("loaded js file..." + jsfile.Filename)

		return true
	}

	return false
}


