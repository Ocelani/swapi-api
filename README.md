
# Americanas SA - Back End Challenge

Código destinado ao desafio Back End da Americanas SA. 

Temática: Star Wars planets API

## Sobre o Projeto

A linguagem de programação utilizada foi **Go (Golang)**, por preferência pessoal.

O projeto segue o [padrão de diretórios recomendado](https://github.com/golang-standards/project-layout) para a linguagem **Go**.

Em relação à arquitetura, foi utilizado os conceitos de [Clean Architecture](https://medium.com/luizalabs/descomplicando-a-clean-architecture-cf4dfc4a1ac6), para abstrair a lógica de negócio, a camada de aplicação, a camada de rede e a camada de dados em módulos distintos, o que possibilita em uma redução do acoplamento no código.

### Observação

O uso de [Protocol Buffers (Protobuf)](https://developers.google.com/protocol-buffers) foi implementado como uma feature, sendo um recurso extra auxiliar.
Os arquivos gerados pelo Protobuf já estão inclusos no projeto, na pasta `./gen`. 
Caso desejado realizar os procedimentos para gerar o código designado ao Protobuf novamente, [siga esse quick start do gRPC](https://www.grpc.io/docs/languages/go/quickstart/).

## Execução

Para executar o projeto, é necessário possuir previamente os seguintes recursos instalados em sua máquina:
- [Golang](https://golang.org) ou [Docker](https://www.docker.com/get-started)

### Execução com Docker

Através do [Makefile](https://pt.wikipedia.org/wiki/Makefile), é possível levantar o container Docker com o comando:
```shell
# com Makefile
make up

# sem Makefile
docker-compose up --build -d 
```

Aguarde o Docker fazer o download das imagens necessárias e inicializar o programa.

Uma vez que a **API** iniciar, ela será **acessível na porta 80** (`localhost:80`).

Na **porta 8080**, encontramos o dashboard do proxy reverso, chamado [Traefik](https://doc.traefik.io/traefik/). O dashboard exibe os serviços que estão em execução no cluster e o endereço de referência de cada um deles.

A aplicação, através do cluster de containers Docker, faz o uso do banco de dados [MongoDB](https://www.mongodb.com/), do qual é populado (por padrão, com a possibilidade de ser desabilitado) com os [dados originais fornecidos previamente pela SWAPI](https://swapi.dev/documentation#planets).

### Execução com Golang

Para executar com a linguagem Go:
```shell
make

# ou

go build -o ./bin ./... # compila o programa
./bin/api               # executa o binário compilado
```

Uma vez que a **API** iniciar, ela será **acessível na porta 80** (`localhost:80`).

#### Flags disponíveis

Por padrão, a **API** espera o uso de um banco de dados MongoDB. No entanto, através de um adapter para repositório de dados, foi implementado a opção de utilizar a própria memória cache como uma base de dados. 

* `-cache-db` *(bool)*: Utilizar a memória cache como banco de dados. 
* `-skip-data `*(bool)*: Não popular o baco de dados com os disponibilizados pela [SWAPI original](https://swapi.dev/documentation#planets).
* `-port` *(int)*: Alternar a porta exposta em que a API executa.

##### Exemplos
```shell
./bin/api -cache-db
./bin/api -skip-data
./bin/api -port 4000
```

#### CLI

Também foi criado uma interface de linha de comando (CLI) para operações de consulta aos dados fornecidos, do qual não necessita da API estar em execução. 

Para o uso de CLI, temos as flags:
- `-id` *(int)*: Pesquisa um respectivo **planet ID**.
- `-name` *(string)*: Pesquisa um respectivo **planet name**.

##### Exemplos
```shell
./bin/cli -id 4
./bin/cli -name Hoth
```

## Testes

É possível realizar **Go tests** com os comandos:
```shell
make test

# testes com MongoDB
make test-mongo     

# testes com os dados disponibilizados
make test-get-data  
```

Para testar os endpoints, foram criados alguns scripts com `curl` em `bash`, disponíveis no diretório `scripts`. Esses scripts podem ser acessados através do Makefile com os seguintes comandos:
```shell
# POST `localhost:80/planet`
make req-create

# GET `localhost:80/planet`
make req-read-all

# GET `localhost:80/planet/1` (id=1)
make req-read-one

# DELETE `localhost:80/planet/1` (id=1)
make req-delete
```

### Obrigado!
*Otávio Celani*
