FROM golang:1.16

WORKDIR /app

COPY ./http-echo-server.go .
COPY ./index.html ./static/index.html

RUN go build -o https-echo-server ./http-echo-server.go

EXPOSE 8080

CMD [ "./https-echo-server" ]
