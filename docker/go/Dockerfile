# Базовый образ
FROM golang:latest
# Папка приложения
ARG APP_DIR=app
# Копирование файлов
COPY . /go/tmp/src/${APP_NAME}
# Рабочая директория
WORKDIR /go/tmp/src/${APP_NAME}

#ENTRYPOINT go build commands/runserver.go
# Библиотеки
# RUN go get <путь_к_библиотеке>
# RUN go get github.com/streadway/amqp
# RUN go get github.com/streadway/amqp