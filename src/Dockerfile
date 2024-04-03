FROM golang:1.22.1-bookworm

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /be-4GLTE-docker

EXPOSE 3321

CMD [ "/be-4GLTE-docker" ]