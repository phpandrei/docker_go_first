# Базовый образ
FROM golang:latest
# Папка приложения
ARG APP_DIR=app
ENV CGO_ENABLED=0
# Копирование файлов
COPY . /go/tmp/src/${APP_NAME}
# Рабочая директория
WORKDIR /go/tmp/src/${APP_NAME}

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
#RUN chmod +x wait-for-postgres.sh

#ENTRYPOINT go build commands/runserver.go
# Библиотеки
# RUN go get <путь_к_библиотеке>
# RUN go get github.com/streadway/amqp
# RUN go get github.com/streadway/amqp