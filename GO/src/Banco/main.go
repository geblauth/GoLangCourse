package main

import (
	"banco/contas"
	"fmt"
)

func main() {

	ContaGermano := contas.ContaCorrente{}
	ContaGermano.Depositar(100000)
	fmt.Println(ContaGermano.VerSaldo())
	pagarBoleto(&ContaGermano, 100)
	fmt.Println(ContaGermano.VerSaldo())

}

func pagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}
