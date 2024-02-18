# docker_go_first

docker-compose exec golang_first go run server.go - запуск сервера. Работает в браузере http://127.0.0.1:8050/test?abc=444444&cde=5555

docker-compose exec golang_first go run hello.go - в консоле

docker-compose exec golang_first gofmt hello.go - потраха файла

docker-compose up --build -d

docker-compose down
 
git pull origin main

git add --all

git commit --all

git push -u origin main

БД
postgres/qwerty

phpandrei

akozlov@akozlov:~/projects/gofirst$ docker-compose exec golang_first go mod init gofirst
go: creating new go.mod: module gofirst
go: to add module requirements and sums:
        go mod tidy

akozlov@akozlov:~/projects/gofirst$ docker-compose exec golang_first go get 'github.com/lib/pq'
go: downloading github.com/lib/pq v1.10.9
go: added github.com/lib/pq v1.10.9
akozlov@akozlov:~/projects/gofirst$ 