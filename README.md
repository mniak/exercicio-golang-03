Exercício 03
===================

## Como rodar

**Warning:** Caso você tenha um posgres rodando na sua máquina, deve parar ele antes de tentar subir o meu banco de dados pelo `docker-compose`, pois o meu também roda na porta `5432`.

### ~~Em um container Docker~~ não tá funcionando

Clone o projeto e num terminal, dentro do diretório onde o repositório for clonado rode as seguintes linhas:
```bash
docker-compose build
docker-compose up
```
O projeto estará rodando em http://localhost:9000

### Na raça (sem containers)

#### No MacOS ou no Linux

**Clona e executa o projeto**
```bash
# Cria uma pasta, seta como GOPATH e coloca a pasta bin no PATH
mkdir ~/gocode
export GOPATH=~/gocode
export PATH="$PATH:$GOPATH/bin"

# Instala o revel
go get github.com/revel/revel
go get github.com/revel/cmd/revel

# Clona o repositorio
git clone https://github.com/mniak/exercicio-golang-03.git $GOPATH/src/exercicio-golang-03
cd $GOPATH/src/exercicio-golang-03

# Roda o banco de dados
docker-compose up -d database

# Executa o projeto
revel run -a exercicio-golang-03
```
O projeto estará rodando em http://localhost:9000

Caso o _revel_ já tenha sido instalado e o projeto já tenha sido clonado usando as instruções acima mas você tenha fechado o terminal e agora não sabe como rodar em um novo terminal, use o trecho abaixo.

```bash
# Seta GOPATH e coloca a pasta bin no PATH
export GOPATH=~/gocode
export PATH="$PATH:$GOPATH/bin"
cd $GOPATH/src/exercicio-golang-03

# Roda o banco de dados
docker-compose up -d database

# Executa o projeto
revel run -a exercicio-golang-03
```
O projeto estará rodando em http://localhost:9000

#### No Windows
Os scripts devem ser rodados no Powershell

**Clona e executa o projeto**
```powershell
# Cria uma pasta, seta como GOPATH e coloca a pasta bin no PATH
mkdir ~/gocode
$env:GOPATH="$env:USERPROFILE/gocode"
$env:PATH="$env:PATH;$env:GOPATH/bin"

# Instala o revel
go get github.com/revel/revel
go get github.com/revel/cmd/revel

# Clona o repositorio
git clone https://github.com/mniak/exercicio-golang-03.git "$env:GOPATH/src/exercicio-golang-03"
cd "$env:GOPATH/src/exercicio-golang-03"

# Roda o banco de dados
docker-compose up -d database

# Executa o projeto
revel run -a exercicio-golang-03
```
O projeto estará rodando em http://localhost:9000

Caso o _revel_ já tenha sido instalado e o projeto já tenha sido clonado usando as instruções acima mas você tenha fechado o terminal e agora não sabe como rodar em um novo terminal, use o trecho abaixo.

Lembre-se de usar o Powershell

```powershell
# Seta GOPATH e coloca a pasta bin no PATH
$env:GOPATH="$env:USERPROFILE/gocode"
$env:PATH="$env:PATH;$env:GOPATH/bin"

# Opcionalmente entra no diretório
cd "$env:GOPATH/src/exercicio-golang-03"

# Roda o banco de dados
docker-compose up -d database

# Executa o projeto
revel run -a exercicio-golang-03
```
O projeto estará rodando em http://localhost:9000


## Credenciais para uso na API
A aplicação vai ler das seguintes variáveis de ambiente

- **`SUPERLINK_CLIENT_ID`:** Client ID
- **`SUPERLINK_CLIENT_SECRET`:** Client Secret

Para rodar a aplicação usando containers via `docker-compose`, deve ser alterado o arquivo `.env` para conter os valores desejados e adequados.
