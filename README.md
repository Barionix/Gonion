
# Gonion
## Gonion é um buscador/verificador de dominios .onion

##Uso
### - Verificador - 
```
./gonion -validate=NOME_DO_ARQUIVO_COM_OS_LINKS  
```  
### - Buscador - 
```
./gonion -range=QUANTIDADE_DE_LINKS_A_SER_TESTADO
```
_PS: Você pode deixar o comando rodando em background, basta usar_ ``` ./gonion -range=QUANTIDADE &``` 

### Funcionamento 
#### Gerador 
_O gerador do Gonion cria links aleatórios na faixa de 22 caracteres e os requisita utilizando o protocolo SOCKS para resolver os enderaços .onion_
#### Verificador  
_O verificador do Gonion lê os arquivos de link e os testa utilizando o protocolo SOCKS para resolver os enderaços .onion_
<<<<<<< HEAD

=======
>>>>>>> a4700f8ffaaa83bf79bb71d0a55ad1201becf9c8
