package utils

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"h12.io/socks"
)

/* Struct de configuração */
type Conf struct {
	Limit   int
	Current int
	Found   int
}

var C Conf

/* Checha a quantidade de routines em execução */
func (c *Conf) Check_rountine() bool {
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
func Check_error(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}

/* Imprime o Banner */
func Banner() {
	banner, err := ioutil.ReadFile("output/banner.txt")
	Check_error(err)
	fmt.Print(string(banner))
}

/* Escreve o log */
func Log(str string) {
	res := Read("log") + str
	btwrite := []byte(res)
	err := ioutil.WriteFile("output/log.txt", btwrite, 0644)
	Check_error(err)
}

/* lê e retorna o conteúdo dos arquivos */
func Read(filename string) string {
	dat, err := ioutil.ReadFile("output/" + filename + ".txt")
	Check_error(err)
	return string(dat) + "\n"
}

/* Escreve o output */
func Write(link string) {
	res := Read("links") + strings.Split(link, " ")[0]
	btwrite := []byte(res)
	err := ioutil.WriteFile("output/links.txt", btwrite, 0644)
	Check_error(err)
}

func Write_Check(link string) {
	res := Read("checked") + strings.Split(link, " ")[0]
	btwrite := []byte(res)
	err := ioutil.WriteFile("output/checked.txt", btwrite, 0644)
	Check_error(err)
}

/* Verifica se o link está no arquivo de output */
func Already_wrote(link string) bool {
	var res bool
	dat, err := ioutil.ReadFile("output/links.txt")
	Check_error(err)
	if strings.Contains(string(dat), link) {
		res = false
	} else {
		res = true
	}
	return res
}

func Already(link string) bool {
	var res bool
	dat, err := ioutil.ReadFile("output/checked.txt")
	Check_error(err)
	if strings.Contains(string(dat), link) {
		res = false
	} else {
		res = true
	}
	return res
}

/* Cria o proxy e retorna o transport */
func Dialer() *http.Transport {
	dialSocksProxy := socks.DialSocksProxy(socks.SOCKS5, "127.0.0.1:9050")
	tr := &http.Transport{Dial: dialSocksProxy}
	return tr
}

/* Checa o link e retorna conteúdo */
func Check_and(link string) string {
	var res string
	tr := Dialer()
	httpClient := &http.Client{Transport: tr}
	fmt.Println("tentando com " + link)
	resp, err := httpClient.Get(link)
	if err != nil {
		/* Cria uma nova conexão*/
		fmt.Println(err)
		res = "ERROR"
		_ = resp
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		Check_error(err)
		res = string(body)
		Log(link + "--> OK")
	}
	return res
}

/* Check o link e retorna uma bool */
func Check(link string) bool {
	var res bool
	tr := Dialer()
	httpClient := &http.Client{Transport: tr}
	_, err := httpClient.Get(link)
	if err != nil {
		Log(link + " --> 404")
	} else {
		res = true
		Log(link + " --> OK")
	}
	return res
}

/* Algoritimo de busca por crawl
h */
func Bar(link string) {
	a := Check_and(link)
	Write_Check(link)
	re, err := regexp.Compile("http://([A-Za-z0-9])*.onion/([A-Za-z0-9])*")
	Check_error(err)
	res := re.FindAllStringSubmatch(a, -1)
	fmt.Println(len(res))
	for _, lk := range res {
		if C.Check_rountine() {
			if Already_wrote(lk[0]) {
				C.Found += 1
				Write(lk[0])
			}
		}
	}
	fmt.Println("[ Done with: " + link + " Found " + strconv.Itoa(C.Found) + " Links ]")
	C.Found = 0
	File_bar("output/links.txt")

}

/* Algoritimo de busca que gera links aleatórios e os testa */
func Generate() {
	var res []string
	var comp string
	var verified bool
	alpha := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 1; i <= 22; i++ {
		res = append(res, alpha[rand.Intn(len(alpha))])
	}
	comp = strings.Join(res, "") + ".onion/"
	verified = Check(comp)
	if verified {
		Write(comp)
	}
	runtime.Goexit()
}

/* Valida links de um arquivo */
func Valid(file string) {
	dat, err := ioutil.ReadFile(file)
	Check_error(err)
	re, err := regexp.Compile(`(.*)/`)
	Check_error(err)
	res := re.FindAllStringSubmatch(string(dat), -1)
	if C.Check_rountine() {
		go func() {
			for _, link := range res {
				a := Check(link[0])
				if a == true {
					Write(link[0])
				}
			}
		}()
	}
}

/* Passa um arquivo de links como parametro para o algoritmo de crawling */
func File_bar(file string) {
	dat, err := ioutil.ReadFile(file)
	Check_error(err)
	re, err := regexp.Compile("http://([A-Za-z0-9])*.onion/([A-Za-z0-9])*")
	Check_error(err)
	res := re.FindAllStringSubmatch(string(dat), -1)
	for _, link := range res {
		if Already(link[0]) {
			Bar(link[0])
		}
	}
}

/* Passa os parametros para o algoritmo que gera aleatório */
func Gen(rang string) {
	ranger, err := strconv.Atoi(rang)
	Check_error(err)
	for i := 0; i <= ranger && C.Check_rountine(); i++ {
		go Generate()
	}
}

/* Faz um parse das opções do usuário */
func Parse(Action string, Limit int) {
	switch Action {
	case "generate":
		if Limit == 50 {
			Gen(os.Args[2])
		} else {
			Gen(os.Args[3])
		}
	case "bar":
		if Limit == 50 {
			if strings.HasPrefix(os.Args[2], ":") {
				File_bar(os.Args[2])
			}
			Bar(os.Args[2])
		} else {
			if strings.HasPrefix(os.Args[3], ":") {
				File_bar(os.Args[3])
			}
			Bar(os.Args[3])
		}
	case "validate":
		if Limit == 50 {
			Valid(os.Args[2])
		} else {
			Valid(os.Args[3])
		}
	default:
		fmt.Println("WHAT THE HECK YOU WANT?")
	}
}
