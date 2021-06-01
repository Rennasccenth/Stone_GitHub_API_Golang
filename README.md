# Stone_GitHub_API_Golang
Teste prático para vaga de back-end Python / Go na Stone.


*Nota 1: Utilizar a branch Release*
*Nota 2: É necessário adicionar o TOKEN de acesso do GITHUB no .ENV da Aplicação.*


## 1 - Como subir a api?
Após clonar o repositório, a estrutura de pastas deve ficar em algo como:

`$GOPATH/src/Stone_GitHub_API_Golang`

Feito isso, basta entrar na pasta raiz do projeto e executar o comando abaixo
para criar a imagem da aplicação:

`docker build . -t stone_backend_rennasccenth`

Com nossa devida imagem criada, podemos executá-la através do comando:

`docker run -d --name=rennasccenth_teste -p 8080:8080 stone_backend_rennasccenth:latest`

Pronto, agora temos a aplicação rodando na porta 8080 local, então certifique-se de que essa porta esteja livre!

## 2 - E quanto aos testes??
Para executar os testes, podemos executar o seguinte comando, ainda na pasta raiz do projeto e com a aplicação rodando:

`godotenv go test ./...`

## 3 - Facilitando os testes...
Para facilitar os testes, deixarei a collection do postman com as requisições que foram utilizadas para validar as features juntamente com o environment usado, estão na pasta:

`$GOPATH/src/Stone_GitHub_API_Golang/postman`

## Aos reviewers!
Essa foi minha primeira experiência com Golang, conhecia a linguagem somente por cima e sabia que lembrava um pouco de C, então aceitei o desafio/sugestão da Lídia de fazer o teste na linguagem que eu menos tinha domínio (nesse caso era 0 mesmo uheuhuaheu).
Dito isso, caso tenham algo a acrescentar ou apontar erros cometidos, sintam-se totalmente a vontade!

## O que foi feito?

- [x] Dockerização
- [x] Documentação
- [x] Testes

### Tier 1
- [x] Retornar Repositório com mais Stars.
- [x] Retornar a Issue aberta que possui mais comentários.
- [x] Retornar os Pull Requests abertos que ainda não foram interagidos.

### Tier 2
- [x] Implementação de uma camada de cache, para as rotas do tier 1.

