package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroCont    int
	saldo         float64
}

func main() {
	contaSoComNome := ContaCorrente{titular: "SomenteNome"}
	contaDoGustavo := ContaCorrente{"Gustavo", 123, 54321, 125.50}
	contaDaLeticia := ContaCorrente{"Leticia", 321, 12345, 152.05}
	fmt.Println(contaSoComNome)
	fmt.Println(contaDoGustavo)
	fmt.Println(contaDaLeticia)

	var contaDaMae *ContaCorrente
	contaDaMae = new(ContaCorrente)
	contaDaMae.titular = "Mae"
	contaDaMae.saldo = 50

	fmt.Println(contaDaMae)
	fmt.Println(*contaDaMae)
}
