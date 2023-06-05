package main

import (
	"fmt"

	"github.com/LuiSnake3/goDesde0/variables"
)

func main() {
estado, texto := variables.ConviertoaTexto(2500)
fmt.Println(estado)
fmt.Println(texto)
}

