package contas

import "github.com/mrcsmotta1/sistemaBancarioPOO/src/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	}

	return "Saldo Insuficiente!"
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0

	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso!", c.saldo
	}

	return "Valor do Deposito menor que 0!", valorDoDeposito
}

func (c *ContaPoupanca) Transferir(valorDeTransferencia float64, contDestino *ContaPoupanca) bool {
	podeTransferir := valorDeTransferencia > 0 && valorDeTransferencia <= c.saldo

	if podeTransferir {
		c.saldo -= valorDeTransferencia
		contDestino.Depositar(valorDeTransferencia)

		return true
	}

	return false
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
