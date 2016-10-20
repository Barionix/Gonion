
# Gonion ##  0.0.2
## Gonion é um buscador/verificador de domínios .onion

##Uso
### - Opções(Flag) -
_-action=_ - Define a ação/metodo a se usar.
```
bar :arquivo_txt.txt
bar link
generate quantidade
```
_-limit=_ - Define o limite máximo de gouroutines.
(Recomendado a partir de 500)
### - Verificador -
```
./gonion -validate=NOME_DO_ARQUIVO_COM_OS_LINKS  
```  
### - Buscador (by RANGE)-
```
./gonion -action=generate QUANTIDADE_DE_LINKS_A_SER_GERADO
```
_PS: Você pode deixar o comando rodando em background, basta usar_
```
./gonion -action=generate QUANTIDADE &
```
### - Buscador (by Crawling)-
Para esse algoritimo você terá de passar um link ou um arquivo de links como parametro.
```
./gonion -action=bar :txt_de_links.txt  
./gonion -action=bar link
```

### Funcionamento
#### Generate
_O gerador do Gonion cria links aleatórios na faixa de 22 caracteres e os requisita utilizando o protocolo SOCKS para resolver os enderaços .onion_
#### Bar
_O algoritimo do Bar usa scraping para buscar no(s) site(s) passado(s) como parametr outros links e usa o mesmo como parametro para uma nova busca. Para usar o Bar é necessário setar o limite de Goroutines, uma quantidade pequena pode resultar num mal funcionamento do programa é recomendado que use algo a partir de 500 Goroutines._
#### Verificador  
_O verificador do Gonion lê os arquivos de link e os testa utilizando o protocolo SOCKS para resolver os enderaços .onion_
