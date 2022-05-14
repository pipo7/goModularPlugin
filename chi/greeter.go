// must be defined in main package .. but note this is chinese plugin
package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("你好宇宙")
}

var Greeter greeting