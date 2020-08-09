package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 2

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
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Gustavo"
	versao := 1.1
	fmt.Println("Olá, mr.", nome, "!")
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("Valor inserido:", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://random-status-code.herokuapp.com", "https://random-status-code.herokuapp.com", "https://random-status-code.herokuapp.com"}

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		fmt.Println("")
		if i != monitoramentos-1 {
			time.Sleep(delay * time.Second)
		}
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso! Status Code:", resp.StatusCode)
	} else {
		fmt.Println("Site:", site, "está fora do ar! Status Code:", resp.StatusCode)
	}
}

func testeDeSlices() {
	teste := []string{}
	fmt.Println("Itens:", teste, ". Length:", len(teste), ". Capacidade:", cap(teste))
	teste = append(teste, "eae")
	fmt.Println("Itens:", teste, ". Length:", len(teste), ". Capacidade:", cap(teste))
	teste = append(teste, "eae")
	fmt.Println("Itens:", teste, ". Length:", len(teste), ". Capacidade:", cap(teste))
	teste = append(teste, "eae")
	fmt.Println("Itens:", teste, ". Length:", len(teste), ". Capacidade:", cap(teste))
	nomes := []string{"https://random-status-code.herokuapp.com", "https://www.alura.com.br", "https://www.caelum.com.br"}
	fmt.Println("Itens:", nomes, ". Length:", len(nomes), ". Capacidade:", cap(nomes))
	nomes = append(nomes, "Novo")
	fmt.Println("Itens:", nomes, ". Length:", len(nomes), ". Capacidade:", cap(nomes))
	nomes = append(nomes, "Novo2", "Novo3")
	fmt.Println("Itens:", nomes, ". Length:", len(nomes), ". Capacidade:", cap(nomes))
	nomes = append(nomes, "Novo4")
	fmt.Println("Itens:", nomes, ". Length:", len(nomes), ". Capacidade:", cap(nomes))
}
