package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroConta, NumeroAgencia int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podesacar := valorDoSaque <= c.saldo && valorDoSaque > 0

	if podesacar {
		c.saldo -= valorDoSaque
		return "Saque Realizado!"
	} else {
		return "saldo insuficiente!"
	}
}

func (c *ContaCorrente) Transferir(valorTransference float64, contaDestino *ContaCorrente) bool {
	if valorTransference < c.saldo && valorTransference > 0 {
		c.saldo -= valorTransference
		contaDestino.Depositar(valorTransference)
		return true

	} else {
		return false
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito Realizado", c.saldo
	} else {
		return "Não depositado", c.saldo
	}
}

func (c *ContaCorrente) VerSaldo() float64 {
	return c.saldo
}
