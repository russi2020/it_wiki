#Версия 1. Размер 26.2MB. Доработать с учетом переменных .env для подключения к базе данных.
FROM golang:alpine as builder

WORKDIR /itWiki

ADD go.mod .

COPY . .

RUN go build -o main .

FROM alpine

WORKDIR /itWiki

COPY --from=builder /itWiki/main /itWiki/main

RUN chmod +x /itWiki/main

CMD ["./main"]

#Версия 2. Размер 9.61MB. Доработать с учетом переменных .env для подключения к базе данных.
#FROM scratch
#ADD wiki /
#CMD ["/wiki"]