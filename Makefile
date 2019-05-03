build:
	cd ./cmd/svc;GOOS=linux GOARCH=amd64 go build  .;
up:
	docker-compose up --build -d;
down:
	docker-compose down;
clean:
	rm ./cmd/svc/svc