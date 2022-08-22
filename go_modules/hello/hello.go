package main

import (
	"fmt"

	"gastonExample.com/greetings"

)

func main() {
	message := greetings.Hello("Gaston")
	fmt.Printf(message)
}