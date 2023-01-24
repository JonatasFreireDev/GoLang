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

const monitoramentos = 5
const delay = 5

func main() {
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("0 - Sair do Programa")
			os.Exit(1)
		}
	}
}

func leComando() int {
	var comandoLido int

	// fmt.Scanf("%d", &comandoLido)
	fmt.Scan(&comandoLido)

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// site := "https://www.alura.com.br"
	// sites := []string{
	// 	"https://www.alura.com.br",
	// 	"https://www.google.com.br",
	// }

	sites, _ := leSitesDoArquivo()

	fmt.Println(sites)

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			response, err := http.Get(site)

			if err != nil {
				fmt.Println("Ocorreu um erro ao realizar o get.", err)
			}

			if response.StatusCode == 200 {
				fmt.Println("Site:", site, "carregado com sucesso!")
				registraLog(site, true)
			} else {
				fmt.Println("Site:", site, "esta com problemas")
				fmt.Println("Status Code:", response.StatusCode)
				registraLog(site, false)
			}
			time.Sleep(delay * time.Second)
		}
	}

}

func leSitesDoArquivo() ([]string, error) {
	var sites []string

	file, err := os.Open("sites.txt")
	defer file.Close()
	// file, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return nil, err
	}

	leitor := bufio.NewReader(file)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Ocorreu um erro:", err)
		}

		sites = append(sites, linha)
	}

	return sites, nil
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer arquivo.Close()

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- Online:" + strconv.FormatBool(status) + "\n")
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
