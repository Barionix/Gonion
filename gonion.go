package main

import (
  "math/rand"
  "strings"
  "h12.me/socks"
  "net/http"
  "os"
  "fmt"
  "flag"
  "io/ioutil"
)
func check_error(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}
func log(str string) {
  res := read("log") + str
  btwrite := []byte(res)
	err := ioutil.WriteFile("log.txt",btwrite, 0644)
  check_error(err)
}
func generate() {
  var res []string
  var comp string
  var verified bool
  alpha := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o","p", "q", "r", "s","t", "u", "v", "w", "x", "y", "z", "1","2","3", "4", "5", "6", "7", "8", "9"}
  for i := 1; i <= 22; i++ {
    res = append(res, alpha[rand.Intn(len(alpha))])
  }
  comp = strings.Join(res, "") + ".onion/"
  verified = check(comp)
  if verified == true {
    log("FUCKING WORKING")
    go write(comp)
    }
}
func check(link string) bool {
  var res bool
  dialSocksProxy := socks.DialSocksProxy(socks.SOCKS5, "127.0.0.1:9050")
  tr := &http.Transport{Dial: dialSocksProxy}
  httpClient := &http.Client{Transport: tr}
  resp, err := httpClient.Get(link)
  if err != nil {
    log(link + "--> 404")
    _ = resp
  } else {
    res = true
    go log(link + "--> OK")
  }
  return res
}
func read(filename string) string {
		dat, err := ioutil.ReadFile(filename + ".txt")
		check_error(err)
		return string(dat) + "\n"
}
func write(link string) {
  res := read("links") + link
  btwrite := []byte(res)
	err := ioutil.WriteFile("links.txt",btwrite, 0644)
  check_error(err)
}
func validate(file string) {
  var tryer []string
  dat, err := ioutil.ReadFile(file)
  check_error(err)
  a := strings.Split(string(dat), "http")
  for _, i := range(a) {
    tryer = append(tryer, "http" + strings.Split(i, " ")[0])
  }
  for _, link := range(tryer) {
    a := check(link)
    if a == true {
      write(link)
    }
  }
}
func main() {
  var ranger = flag.Int("range", 0, "Configuração de range.")
  var file = flag.String("validate", "0", "Arquivo para validação.")
  flag.Parse()
  if *ranger != 0 {
    for qntd := 0; qntd <= *ranger; qntd++ {
      generate()
    }
  } else {
    if *file != "0" {
      validate(*file)
    }
  }
  fmt.Println("[FINISHED]")
}
