package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 3

//Tempo que será esperado para a próxima verificação de sites
const tempoDeEspera = 5 // segundos

const localArquivo = "sites.txt"

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "José"
	versao := 1.2
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi: ", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	//site := "https://www.alura.com.br"
	// sites := []string{"https://random-status-code.herokuapp.com", "https://ww.alura.com.br",
	// 	"https://caelum.com.br"}

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		fmt.Println("Aguarde...")
		time.Sleep(tempoDeEspera * time.Second)
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro ao testar site! Erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code: ",
			resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open(localArquivo)
	// arquivo, err := ioutil.ReadFile(localArquivo)

	if err != nil {
		fmt.Println("Problemas ao abrir o arquivo ", localArquivo, "! Erro:", err)
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
