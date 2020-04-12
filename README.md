# Gonion #Versão#  0.0.2


Aplicação desenvolvida na linguagem Go com a função de buscar e/ou verificar domínios  *.onion*

## Uso
### Opções(Flag)

Opção               |  Descrição              
------------------- |  ---------
`-action=`          | define a ação/método a ser utilizado
`-limit=`           | define o limite máximo de gouroutines*

**limite recomendado a partir de 500**

* Exemplos usando `-action=`:
_-action=generate 500000_ - Gera 50.000 links   
_-action=bar https://zqktlwi4fecvo6ri.onion_ -  Começa a buscar a partir deste link
_-action=link :arquivo.txt_ - Começa a buscar a partir de links em um arquivo.txt

* Exemplos usando `-limit=`:
_-limit=500_ Seta 500 como limite de Goroutines

### Verificador

Uso : `./gonion -validate=<nome do arquivo com os links>`  

### Buscador (by RANGE)

Uso : `./gonion -action=generate <quantidade de links a ser gerada>`

**Obs: É recomendado deixar o comando rodando em background, para isso usar** `&` :

Uso : `./gonion -action=generate <quantidade de links a ser gerada> &`

###Buscador (by Crawling)
Para esse algoritmo você terá de passar um link ou um arquivo de links como parametro.

## Funcionamento
### Generate
_O gerador do Gonion cria links aleatórios na faixa de 22 caracteres e os requisita utilizando o protocolo SOCKS para resolver os enderaços .onion_
### Bar
_O algoritmo do Bar usa scraping para buscar no(s) site(s) passado(s) como parametr outros links e usa o mesmo como parametro para uma nova busca. Para usar o Bar é necessário setar o limite de Goroutines, uma quantidade pequena pode resultar num mal funcionamento do programa é recomendado que use algo a partir de 500 Goroutines._
### Verificador  
_O verificador do Gonion lê os arquivos de link e os testa utilizando o protocolo SOCKS para resolver os enderaços .onion_
