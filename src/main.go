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

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	// c.saldo += c.saldo
	podeDepositar := valorDoDeposito > 0

	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso!", c.saldo
	}

	return "Valor do Deposito menor que 0!", valorDoDeposito
}

func main() {
	contaDoMarcos := ContaCorrente{}
	contaDoMarcos.titular = "Marcos"
	contaDoMarcos.saldo = 500

	fmt.Println(contaDoMarcos.saldo)
	fmt.Println(contaDoMarcos.Sacar(200))
	fmt.Println(contaDoMarcos.saldo)

	fmt.Println(contaDoMarcos.Depositar(100))
	fmt.Println(contaDoMarcos.saldo)
}
