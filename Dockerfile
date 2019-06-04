FROM golang

WORKDIR /app/macaddress_io

COPY macaddress_cli.go .

RUN go build macaddress_cli.go

ENV API_KEY "<insert api_key from macadress.io>"
