package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podesacar := valorDoSaque <= c.saldo && valorDoSaque > 0

	if podesacar {
		c.saldo -= valorDoSaque
		return "Saque Realizado!"
	} else {
		return "saldo insuficiente!"
	}
}
func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito Realizado", c.saldo
	} else {
		return "Não depositado", c.saldo
	}
}
func (c *ContaPoupanca) verSaldo() float64 {
	return c.saldo
}
