# 配置文件目录的绝对路径
config = ~/GolandProjects/go_service_work/config

r:
	go run main.go
b:
	go build main.go
s:
	go run main.go server
test:
	go test -v ./... 2>&1 | go-junit-report > report.xml
coverage:
	gocov test ./... | gocov-xml > coverage.xml
docker:
	docker build -t go-service:dev .
up:
	docker-compose up -d
restart:
	docker-compose restart api
down:
	docker-compose down
server:
	docker run -d --rm --name go-service-api -p 8080:8080 -v $(config):/app/config -e WORK=server -e ENV=config go-service:dev
