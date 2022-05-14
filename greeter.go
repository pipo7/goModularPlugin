package main

import (
	"fmt"
	"os"
	"plugin" // use Go's plugin package
)

type GreeterInMain interface {
	Greet()
}

func main() {
	// determine plugin to load
	lang := "english" // using english as default.
	// OR use os.Args as what is passed via go run greeter.go <lang>
	if len(os.Args) == 2 {
		lang = os.Args[1]
	}
	var mod string
	switch lang {
	case "english":
		mod = "./eng/eng.so"
	case "chinese":
		mod = "./chi/chi.so"
	default:
		fmt.Println("don't speak that language")
		os.Exit(1)
	}

	// load module
	// 1. open the .so plugin file to load the symbols
	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Lookup the exported symbol "Greeter" with plguin.Lookup("Greeter"). Note the symbol name matches the name of the exported package variable defined in the plugin module
	// 2. look up a symbol (an exported function or variable)
	// in this case, variable Greeter
	symGreeter, err := plug.Lookup("Greeter")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type GreeterInMain (defined above) which should implement Greet method
	// and symGreeter from plugin is Greeter which is receiver to Greet method.
	var greeterInMain GreeterInMain
	greeterInMain, ok := symGreeter.(GreeterInMain)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	// 4. use the module. Calls the Greet() method to display the message from the plugin that is dynamically loaded here.
	greeterInMain.Greet()

}
