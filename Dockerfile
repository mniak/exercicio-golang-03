FROM golang:latest
ENV GOPATH=/go/
 
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel

RUN mkdir -p $GOPATH/exercicio-golang-03
COPY . /go/src/exercicio-golang-03
RUN revel build -m prod exercicio-golang-03 /app
 
WORKDIR /app
ENTRYPOINT ./run.sh
EXPOSE 9000