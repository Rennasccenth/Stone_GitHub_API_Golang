FROM amd64/golang:1.16

ENV GO111MODULE=on

WORKDIR /go/src/Stone_GitHub_API_Golang
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["Stone_GitHub_API_Golang"]