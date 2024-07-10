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

func (c *ContaCorrente) Transferir(valorDeTransferencia float64, contDestino *ContaCorrente) bool {
	podeTransferir := valorDeTransferencia > 0 && valorDeTransferencia <= c.saldo

	if podeTransferir {
		c.saldo -= valorDeTransferencia
		contDestino.Depositar(valorDeTransferencia)

		return true
	}

	return false

}

func main() {
	contaDoMarcos := ContaCorrente{}
	contaDoMarcos.titular = "Marcos"
	contaDoMarcos.saldo = 500

	contaDoGuilherme := ContaCorrente{}
	contaDoGuilherme.titular = "Guilherme"
	contaDoGuilherme.saldo = 1000

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
