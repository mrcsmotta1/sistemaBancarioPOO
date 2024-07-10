package main

import (
	"fmt"

	"github.com/mrcsmotta1/sistemaBancarioPOO/src/contas"
)

func main() {
	contaDoMarcos := contas.ContaCorrente{}
	contaDoMarcos.Titular = "Marcos"
	contaDoMarcos.Saldo = 500

	contaDoGuilherme := contas.ContaCorrente{}
	contaDoGuilherme.Titular = "Guilherme"
	contaDoGuilherme.Saldo = 1000

	fmt.Println(contaDoMarcos)
	fmt.Println(contaDoGuilherme)

	status := contaDoGuilherme.Transferir(1000, &contaDoMarcos)
	fmt.Println(status)
	fmt.Println(contaDoMarcos)
	fmt.Println(contaDoGuilherme)
	// fmt.Println(contaDoMarcos.Sacar(200))
	// fmt.Println(contaDoMarcos.saldo)

	// fmt.Println(contaDoMarcos.Depositar(100))
	// fmt.Println(contaDoMarcos.saldo)
}
