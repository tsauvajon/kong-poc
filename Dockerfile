FROM golang:1.16.2-buster

COPY hello.go .

RUN go build -o api hello.go

EXPOSE 8080
CMD [ "./api" ]
