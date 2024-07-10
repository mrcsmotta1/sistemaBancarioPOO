package contas

import "github.com/mrcsmotta1/sistemaBancarioPOO/src/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo

	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	}

	return "Saldo Insuficiente!"
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0

	if podeDepositar {
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso!", c.Saldo
	}

	return "Valor do Deposito menor que 0!", valorDoDeposito
}

func (c *ContaCorrente) Transferir(valorDeTransferencia float64, contDestino *ContaCorrente) bool {
	podeTransferir := valorDeTransferencia > 0 && valorDeTransferencia <= c.Saldo

	if podeTransferir {
		c.Saldo -= valorDeTransferencia
		contDestino.Depositar(valorDeTransferencia)

		return true
	}

	return false
}
