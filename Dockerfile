FROM golang:1.21.5-alpine

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o server . 

CMD [ "/app/server" ]
