package main

import (
	"fmt"

	"github.com/PAJARON/experiments/mod2"
)

func init() {
	fmt.Println("Ready to launch ...")
}

func main() {
	saludo := "Hola!"
	fmt.Println(saludo)
	fmt.Println(mod2.Translate(saludo))
}
