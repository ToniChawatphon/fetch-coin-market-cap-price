FROM golang:1.16.3-alpine3.12

ENV GO111MODULE=on
EXPOSE 8080

WORKDIR /fetch-coin-market-cap-price
COPY go.mod .
COPY go.sum . 

RUN go mod download 
COPY . .
RUN go build 

ENTRYPOINT ["go"]
CMD [ "run", "main.go" ]
