Exemplo 03
===================

## Como rodar
### MacOS/Linux

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
git clone https://github.com/mniak/exercicio-golang-03.git -o $GOPATH/src/exercicio-golang-03

# Executa o projeto
revel run -a desafio-curso-03
```

Caso o _revel_ já tenha sido instalado e o projeto já tenha sido clonado usando as instruções acima mas você tenha fechado o terminal e agora não sabe como rodar em um novo terminal, use o trecho abaixo.

```bash
# Seta GOPATH e coloca a pasta bin no PATH
export GOPATH=~/gocode
export PATH="$PATH:$GOPATH/bin"

# Executa o projeto
revel run -a desafio-curso-03
```

### Windows
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
git clone https://github.com/mniak/exercicio-golang-03.git -o "$env:GOPATH/src/exercicio-golang-03"

# Executa o projeto
revel run -a desafio-curso-03
```

Caso o _revel_ já tenha sido instalado e o projeto já tenha sido clonado usando as instruções acima mas você tenha fechado o terminal e agora não sabe como rodar em um novo terminal, use o trecho abaixo.

Lembre-se de usar o Powershell

```powershell
# Seta GOPATH e coloca a pasta bin no PATH
$env:GOPATH="~/gocode"
$env:PATH="$env:PATH;$env:GOPATH/bin"

# Executa o projeto
revel run -a desafio-curso-03
```
