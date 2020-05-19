package mod2

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("Módulo 2 activado , ...")
	rand.Seed(time.Now().Unix())
}

//Commuta caracteres en una cadena
func Translate(cadena string) string {
	runes := []rune(cadena)
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}
