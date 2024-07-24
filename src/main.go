package main

import (
	"fmt"

	"github.com/mrcsmotta1/sistemaBancarioPOO/src/contas"
)

func main() {
	contaPoupandaMarcos := contas.ContaPoupanca{}
	contaPoupandaMarcos.Depositar(100)
	contaPoupandaMarcos.Sacar(50)

	fmt.Println(contaPoupandaMarcos.ObterSaldo())
}
