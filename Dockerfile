FROM golang:1.19

WORKDIR /usr/src/app
COPY . .

ENV TZ 'Europe/London'

RUN go build
CMD ["./radiotext"]