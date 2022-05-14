// package must be defined as main .. but note this is chinese plugin
package main

import (
	"fmt"
)

type greeting string

func (g greeting) Greet() {
	fmt.Println("Hello Universe")
}

// exported as symbol named "Greeter"
var Greeter greeting

// you can see that Greet method has receiver of type greeting i.e. a string
/* func main() {
	Greeter.Greet()   // Greeter is receiver to method Greet.
	OR Greet method in this english plugin can be called by Greeter, which is exported symbol ...
	fmt.Println(reflect.TypeOf(Greeter))
} */
