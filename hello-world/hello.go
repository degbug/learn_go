package main

import "fmt"

func main() {
	fmt.Println(Hello("degbug"))
}

func Hello(name string) string {
	return "Hello, " + name
}
