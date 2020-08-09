package main

import (
	"fmt"
	"os"
)

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
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com"
	sites[1] = "https://random-status-code.herokuapp.com"
	sites[2] = "https://random-status-code.herokuapp.com"
	sites[3] = "https://random-status-code.herokuapp.com"
	fmt.Println(sites)
	// for sites {
	// 	resp, _ := http.Get(site)

	// 	if resp.StatusCode == 200 {
	// 		fmt.Println("Site:", site, "foi carregado com sucesso!")
	// 	} else {
	// 		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
	// 	}
	// }
}
