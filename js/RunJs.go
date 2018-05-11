package js

import (
	"github.com/robertkrimen/otto"
	"fmt"
)

func RunJs(jsfile *JsFile, params map[string]string) {

	vm := otto.New()

	for name, value := range(params) {
		vm.Set(name, value)
	}



	_, er := vm.Run(jsfile.JsScript)

	if er != nil {
		fmt.Println(er)
	}
}
