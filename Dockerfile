# Контейнер с Lunux, чтобы компилить код под Linux 
FROM golang:1.23.7-alpine3.21 AS bilder
WORKDIR /appf
COPY ./main.go ./
COPY ./handlers.go ./
COPY go.mod ./

RUN go mod tidy
RUN go build -o app .

# Старт кода внутри контейнера
FROM alpine:3.21
WORKDIR /myapp
COPY --from=bilder /appf/app ./
EXPOSE 6080
ENTRYPOINT ["/myapp/app"]