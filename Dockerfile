FROM golang:alpine3.19

WORKDIR /calc

COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o run

CMD [ "./run" ]
