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

const monitoramento = 2
const intervalo = 2

func main() {
	for {
		exibeIntroducao()
		exibeMenu()
		var comando int = leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Logs ")
			imprimeLogs()
		case 3:
			fmt.Println("exit ")
			os.Exit(0)
		default:
			fmt.Println("Não conheço ")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	name := "Fernando"
	version := 44

	fmt.Println("Hello mr.", name)
	fmt.Println("The version of progam is", version)
}

func exibeMenu() {
	fmt.Println("1- Start")
	fmt.Println("2- show logs")
	fmt.Println("0- exit program")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func iniciarMonitoramento() {
	fmt.Println("monitoring... ")

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println(i, ":", site)
			testaSite(site)
		}
		fmt.Println("------")

		time.Sleep(intervalo * time.Second)
	}
}

func testaSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problema, status:", response.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
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

	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("logs.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	fmt.Println(string(arquivo))
}
