package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	introducao()

	for {
		menu()
		comando := leComando()

		switch comando {
		case 1:
			monitoramento()

		case 2:
			fmt.Println("Exibindo...")
			imprimeLogs()

		case 3:
			saiDoPrograma()

		default:
			fmt.Println("Comando não aceito")
			os.Exit(-1)

		}
	}

}

func introducao() {
	var nome string = "Germano"
	var idade = 27

	fmt.Println("Hello World,", nome+" ", idade)
}
func menu() {
	fmt.Println("1 - Monitor")
	fmt.Println("2 - Exibe log")
	fmt.Println("3 - Sair")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)

	return comando
}
func saiDoPrograma() {
	fmt.Println("Saindo...")
	os.Exit(0)
}

func monitoramento() {
	sites := leSites()
	fmt.Println("Monitorando...")

	for i := 0; 1 < monitoramentos; i++ {

		for i, site := range sites {

			fmt.Println("Posição: ", i, "site: ", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}

}
func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Deu Erro!", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site está funcionando")
		registraLog(site, true)
	} else {
		fmt.Println("Site down: ", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSites() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Deu Erro!", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}

	}
	return sites

}
func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()

}
func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}
