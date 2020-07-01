package main

import (
	"flag"
	"fmt"
	"utils"
)

/* Struct de configuração */
type Conf struct {
	Limit   int
	Current int
}

/* Iniciando uma variavel global do tipo Conf */

func main() {
	utils.Banner()
	var Action = flag.String("action", "0", "Configuração de range.")
	var Limit = flag.Int("limit", 500, "Limite de GoRoutines")
	flag.Parse()
	fmt.Println(*Limit)
	//var myip  = utils.Check_and("https://torguard.net/whats-my-ip.php")
	//fmt.Println(myip)
	utils.C = utils.Conf{*Limit, 0, 0}
	utils.Parse(*Action, utils.C.Limit)
	var input string
	fmt.Scanln(&input)
}
