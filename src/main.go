package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	}

	return "Saldo Insuficiente!"
}

func main() {
	contaDoMarcos := ContaCorrente{}
	contaDoMarcos.titular = "Marcos"
	contaDoMarcos.saldo = 500

	fmt.Println(contaDoMarcos.saldo)
	fmt.Println(contaDoMarcos.Sacar(200))
	fmt.Println(contaDoMarcos.saldo)
}
