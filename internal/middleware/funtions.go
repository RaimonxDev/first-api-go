package middleware

import "fmt"

func Greeter(name string) {
	fmt.Println("Hello", name)
}

func Bye(name string) {
	fmt.Println("Bye", name)
}
