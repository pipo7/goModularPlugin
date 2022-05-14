# goModularPlugin
# Build loosely coupled modular programs using packages compiled as shared object libraries that can be loaded and bound to dynamically at runtime.
# A Go plugin is a package compiled using the -buildmode=plugin build flag to produce a shared object (.so) library file. The exported functions and variables, in the Go package, are exposed as ELF symbols that can be looked up and be bound to at runtime using the plugin package
# The Go compiler is capable of creating C-style dynamic shared libraries using build flag -buildmode=c-shared 

# Compile project
# This is will create the .so files
go build -buildmode=plugin -o eng/eng.so eng/greeter.go
go build -buildmode=plugin -o chi/chi.so chi/greeter.go

# Plugins are loaded dynamically using Go’s plugin package. The driver (or client) program ./greeter.go makes use of the plugins, compiled earlier,
# run as 
 go run greeter.go english
Hello Universe
> go run greeter.go chinese
你好宇宙

# references
https://medium.com/learning-the-go-programming-language/writing-modular-go-programs-with-plugins-ec46381ee1a9
https://github.com/vladimirvivien/go-plugin-example/tree/master/chi