package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Gustavo"
	idade := 25
	versao := 1.1
	fmt.Println("Olá, mr.", nome, "! Sua idade é", idade, "anos")
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel idade é", reflect.TypeOf(idade))
	fmt.Println("O tipo da variavel versao é", reflect.TypeOf(versao))
}
