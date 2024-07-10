package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoMarcos := ContaCorrente{}
	contaDoMarcos.titular = "Marcos"
	contaDoMarcos.saldo = 500

	fmt.Println(contaDoMarcos.saldo)
}
