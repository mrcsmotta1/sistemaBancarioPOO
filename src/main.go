package main

import (
	"fmt"

	"github.com/mrcsmotta1/sistemaBancarioPOO/src/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaPoupancaMarcos := contas.ContaPoupanca{}
	contaPoupancaMarcos.Depositar(100)
	PagarBoleto(&contaPoupancaMarcos, 60)

	fmt.Println(contaPoupancaMarcos.ObterSaldo())

	contaCorrenteMarcos := contas.ContaCorrente{}
	contaCorrenteMarcos.Depositar(200)
	PagarBoleto(&contaCorrenteMarcos, -10)

	fmt.Println(contaCorrenteMarcos.ObterSaldo())
}
