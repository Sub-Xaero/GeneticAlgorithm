package ga

import "fmt"

var PrintToConsole = func(a ...interface{}) {
	fmt.Println(a...)
}

var Output = PrintToConsole

func SetOutputFunc(f func(a ...interface{})) {
	Output = f
}
