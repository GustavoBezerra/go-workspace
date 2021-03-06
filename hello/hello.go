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

const monitoramentos = 2
const delay = 1

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
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

	sites := leSitesDoArquivo()

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
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso! Status Code:", resp.StatusCode)
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está fora do ar! Status Code:", resp.StatusCode)
		registraLog(site, false)
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

func leSitesDoArquivo() []string {
	var sites []string

	// arquivo, err := os.Open("sites.txt")
	// arquivo, err := ioutil.ReadFile("sites.txt")
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
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

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString("[" + time.Now().Format(time.Stamp) + "] " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
