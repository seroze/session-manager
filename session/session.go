package session

import "fmt"

func SayHello(name string) {
	fmt.Println("Hello: " + name)
}

func Init() {
	fmt.Println("Session initialized")
}
