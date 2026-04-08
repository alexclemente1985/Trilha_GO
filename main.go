package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	//"reflect"

	"net/http"
	"os"
	"time"
)

const testeConstante = "teste de constante"
const delay = 5

const (
	_                        = iota
	stdLongMonth             = iota + stdNeedDate  // "January"
	stdMonth                                       // "Jan"
	stdNumMonth                                    // "1"
	stdZeroMonth                                   // "01"
	stdLongWeekDay                                 // "Monday"
	stdWeekDay                                     // "Mon"
	stdDay                                         // "2"
	stdUnderDay                                    // "_2"
	stdZeroDay                                     // "02"
	stdHour                  = iota + stdNeedClock // "15"
	stdHour12                                      // "3"
	stdZeroHour12                                  // "03"
	stdMinute                                      // "4"
	stdZeroMinute                                  // "04"
	stdSecond                                      // "5"
	stdZeroSecond                                  // "05"
	stdLongYear              = iota + stdNeedDate  // "2006"
	stdYear                                        // "06"
	stdPM                    = iota + stdNeedClock // "PM"
	stdpm                                          // "pm"
	stdTZ                    = iota                // "MST"
	stdISO8601TZ                                   // "Z0700"  // prints Z for UTC
	stdISO8601SecondsTZ                            // "Z070000"
	stdISO8601ShortTZ                              // "Z07"
	stdISO8601ColonTZ                              // "Z07:00" // prints Z for UTC
	stdISO8601ColonSecondsTZ                       // "Z07:00:00"
	stdNumTZ                                       // "-0700"  // always numeric
	stdNumSecondsTz                                // "-070000"
	stdNumShortTZ                                  // "-07"    // always numeric
	stdNumColonTZ                                  // "-07:00" // always numeric
	stdNumColonSecondsTZ                           // "-07:00:00"
	stdFracSecond0                                 // ".0", ".00", ... , trailing zeros included
	stdFracSecond9                                 // ".9", ".99", ..., trailing zeros omitted

	stdNeedDate  = 1 << 8             // need month, day, year
	stdNeedClock = 2 << 8             // need hour, minute, second
	stdArgShift  = 16                 // extra argument in high bits, above low stdArgShift
	stdMask      = 1<<stdArgShift - 1 // mask out argument
)

func main() {
	exibeNomes()
	exibeIntro()

	for {
		exibeMenu()
		comando := leituraComando()
		switch comando {
		case 1:
			monitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não sei do que vc tá falando...")
			os.Exit(-1)
		}
	}
	// var nome string = "Raul"
	// //var versao float32 = 1.1
	// var versao = 1.1
	// var idade = 24
	// teste := "Isto é uma variável para teste de declaração curta"

	// fmt.Println("Ola sr. ", nome, " vc tem ", idade, " anos")
	// fmt.Println("Programa na versão ", versao)
	// fmt.Println("Tipo da variavel idade: ", reflect.TypeOf(idade))
	// fmt.Println("Tipo da variavel versao: ", reflect.TypeOf(versao))

	// fmt.Println("Tipo da variável teste: ", reflect.TypeOf(teste))
	// fmt.Println(teste)

	// fmt.Println("1- Iniciar Monitoramento")
	// fmt.Println("2- Exibir Logs")
	// fmt.Println("0- Sair do Programa")

	// var comando int
	// //fmt.Scanf("%d", &comando) //precisa que seja inferido o tipo após o %
	// fmt.Scan(&comando) //Inferência automática

	// fmt.Println("O comando escolhido foi o ", comando)

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do programa...")
	// } else {
	// 	fmt.Println("Não conheço este comando")
	// }

	//Não precisa de break pq apenas um case será executado, sem falldown
}

func exibeIntro() {
	nome, idade := devolveNomeEIdade()
	versao := 1.2
	fmt.Println("Olá sr. ", nome)
	fmt.Println("Sua idade: ", idade)
	fmt.Println("Esta é a versão ", versao, " do programa")
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leituraComando() int {
	var comando int
	//fmt.Scanf("%d", &comando) //precisa que seja inferido o tipo após o %
	fmt.Scan(&comando) //Inferência automática

	fmt.Println("O comando escolhido foi o ", comando)

	return comando
}

func monitoramento() {
	fmt.Println("Monitorando...")
	// var sites [4]string //todo array no Go tem que ter tamanho fixo

	// sites[0] = "https://httpbin.org/status/404"
	// sites[1] = "https://www.alura.com.br"
	// sites[2] = "https://www.caelum.com.br"
	// sites := []string{"https://random-status-code.herokuapp.com/",
	// 	"https://www.alura.com.br", "https://www.caelum.com.br"}
	sites := leSitesDoArquivo()

	// for i := 0; i < len(sites); i++ { //Go somente possui FOR como estrutura de loop
	// 	resp, _ := http.Get(sites[i])
	// 	if resp.StatusCode == 200 {
	// 		fmt.Println("Site: ", sites[i], "foi carregado com sucesso!")
	// 	} else {
	// 		fmt.Println("Site: ", sites[i], " está com problema de carregamento. Status Code: ", resp.StatusCode)
	// 	}

	// }

	for i, site := range sites { //Go somente possui FOR como estrutura de loop
		testaSite(site, i)
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

func testaSite(site string, i int) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site ", i+1, ": ", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site ", i+1, ":", site, " está com problema de carregamento. Status Code: ", resp.StatusCode)
		registraLog(site, false)
	}
}

func devolveNomeEIdade() (string, int) {
	nome := "Alex"
	idade := 40
	return nome, idade
}

func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"} //slice: array onde não se define um tamanho inicial
	fmt.Println(len(nomes))
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Aparecida") //adiciona elementos no slice
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens") //sempre que um slice "estoura" sua capacidade atual, ele dobra de tamanho

	fmt.Println(testeConstante)
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}
	//fmt.Println(string(arquivo))
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n') //byte é representado com aspas simples
		// if err != nil {
		// 	fmt.Println("Ocorreu um erro: ", err)
		// }
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sites

}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(arquivo)

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n") //o padrão da formatação são as constantes informadas anteriormente, no local das posições da data e do horário

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	// Definindo a rota "/"
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Backend do Kanban em Go está ONLINE!")
// 	})

// 	fmt.Println("Servidor iniciado em http://localhost:8080")

// 	// Iniciando o servidor na porta 8080
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		fmt.Printf("Erro ao subir o servidor: %v\n", err)
// 	}
// }
