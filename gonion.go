package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"h12.me/socks"
)

/* Struct de configuração */
type Conf struct {
	Limit   int
	Current int
}

/* Checha a quantidade de routines em execução */
func (c *Conf) check_rountine() bool {
	var res bool
	var arr []runtime.StackRecord
	v, _ := runtime.GoroutineProfile(arr)
	c.Current = v
	if c.Current <= c.Limit {
		res = true
	}
	return res
}

/* Evitando i err != nil (hehe)*/
func check_error(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}

/* Iniciando uma variavel global do tipo Conf */
var C Conf

/* Imprime o Banner */
func Banner() {
	banner, err := ioutil.ReadFile("output/banner.txt")
	check_error(err)
	fmt.Print(string(banner))
}

/* Escreve o log */
func log(str string) {
	res := read("log") + str
	btwrite := []byte(res)
	err := ioutil.WriteFile("output/log.txt", btwrite, 0644)
	check_error(err)
}

/* lê e retorna o conteúdo dos arquivos */
func read(filename string) string {
	dat, err := ioutil.ReadFile("output/" + filename + ".txt")
	check_error(err)
	return string(dat) + "\n"
}

/* Escreve o output */
func write(link string) {
	res := read("output/links") + strings.Split(link, " ")[0]
	btwrite := []byte(res)
	err := ioutil.WriteFile("output/links.txt", btwrite, 0644)
	check_error(err)
}

/* Verifica se o link está no arquivo de output */
func already(link string) bool {
	var res bool
	dat, err := ioutil.ReadFile("output/links.txt")
	check_error(err)
	if strings.Contains(string(dat), link) {
		res = false
	} else {
		res = true
	}
	return res
}

/* Cria o proxy e retorna o transport */
func dialer() *http.Transport {
	dialSocksProxy := socks.DialSocksProxy(socks.SOCKS5, "127.0.0.1:9050")
	tr := &http.Transport{Dial: dialSocksProxy}
	return tr
}

/* Checa o link e retorna conteúdo */
func check_and(link string) string {
	var res string
	tr := dialer()
	httpClient := &http.Client{Transport: tr}
	resp, err := httpClient.Get(link)
	if err != nil {
		res = "ERROR"
		_ = resp
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		check_error(err)
		res = string(body)
		log(link + "--> OK")
	}
	return res
}

/* Check o link e retorna uma bool */
func check(link string) bool {
	var res bool
	tr := dialer()
	httpClient := &http.Client{Transport: tr}
	_, err := httpClient.Get(link)
	if err != nil {
		log(link + "--> 404")
	} else {
		res = true
		log(link + "--> OK")
	}
	return res
}

/* Algoritimo de busca por crawling */
func bar(link string) {
	a := check_and(link)
	re, err := regexp.Compile(`http://(.*).onion/`)
	check_error(err)
	res := re.FindAllStringSubmatch(a, -1)
	for _, lk := range res {
		if len(lk) > 0 && C.check_rountine() {
			if already(lk[0]) {
				write(lk[0])
				go bar(lk[0])
			}
		}
	}
	runtime.Goexit()
}

/* Algoritimo de busca que gera links aleatórios e os testa */
func generate() {
	var res []string
	var comp string
	var verified bool
	alpha := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 1; i <= 22; i++ {
		res = append(res, alpha[rand.Intn(len(alpha))])
	}
	comp = strings.Join(res, "") + ".onion/"
	verified = check(comp)
	if verified {
		log("FUCKING WORKING")
		write(comp)
	}
	runtime.Goexit()
}

/* Valida links de um arquivo */
func valid(file string) {
	dat, err := ioutil.ReadFile(file)
	check_error(err)
	re, err := regexp.Compile(`(.*)/`)
	check_error(err)
	res := re.FindAllStringSubmatch(string(dat), -1)
	if C.check_rountine() {
		go func() {
			for _, link := range res {
				a := check(link[0])
				if a == true {
					write(link[0])
				}
			}
		}()
	}
}

/* Passa um arquivo de links como parametro para o algoritmo de crawling */
func file_bar(file string) {
	dat, err := ioutil.ReadFile(strings.Split(file, ":")[1])
	check_error(err)
	re, err := regexp.Compile(`(.*)/`)
	check_error(err)
	res := re.FindAllStringSubmatch(string(dat), -1)
	for _, link := range res {
		if C.check_rountine() {
			go bar(link[0])
		}
	}
}

/* Passa os parametros para o algoritmo que gera aleatório */
func gen(rang string) {
	ranger, err := strconv.Atoi(rang)
	check_error(err)
	for i := 0; i <= ranger && C.check_rountine(); i++ {
		go generate()
	}
}

/* Faz um parse das opções do usuário */
func parse(Action string, Limit int) {
	switch Action {
	case "generate":
		if Limit == 50 {
			gen(os.Args[2])
		} else {
			gen(os.Args[3])
		}
	case "bar":
		if Limit == 50 {
			if strings.HasPrefix(os.Args[2], ":") {
				file_bar(os.Args[2])
			}
			bar(os.Args[2])
		} else {
			if strings.HasPrefix(os.Args[3], ":") {
				file_bar(os.Args[3])
			}
			bar(os.Args[3])
		}
	case "validate":
		if Limit == 50 {
			valid(os.Args[2])
		} else {
			valid(os.Args[3])
		}
	default:
		fmt.Println("WHAT THE HECK YOU WANT?")
	}
}

func main() {
	Banner()
	var Action = flag.String("action", "0", "Configuração de range.")
	var Limit = flag.Int("limit", 50, "Limite de GoRoutines")
	flag.Parse()
	C = Conf{*Limit, 0}
	parse(*Action, C.Limit)
	var input string
	fmt.Scanln(&input)
}
