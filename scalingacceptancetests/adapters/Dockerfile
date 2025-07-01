FROM golang:1.24-alpine

ARG bin_to_build
ARG exposed_port

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o svr scalingacceptancetests/cmd/${bin_to_build}/main.go

EXPOSE ${exposed_port}

CMD [ "./svr" ]