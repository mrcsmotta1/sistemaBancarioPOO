package main

import (
	"fmt"

	"github.com/mrcsmotta1/sistemaBancarioPOO/src/clientes"
	"github.com/mrcsmotta1/sistemaBancarioPOO/src/contas"
)

func main() {
	clienteMarcos := clientes.Titular{
		Nome:      "Marcos",
		CPF:       "308.916.774-61",
		Profissao: "Desenvolvedor",
	}

	contaDoMarcos := contas.ContaCorrente{Titular: clienteMarcos,
		NumeroAgencia: 123,
		NumeroConta:   123456,
		Saldo:         100,
	}

	fmt.Println(contaDoMarcos)

}
