docker-save:
	docker save src-app -o artifacts/src.tar

docker-load:
	docker load -i artifacts/src.tar

docker-run:
	docker compose up -d --build

docker-stop:
	docker compose down

app:
	go build -o ./artifacts/main.exe ./cmd/main.go
	./artifacts/main.exe

clean:
	rmdir /Q /S artifacts
	mkdir artifacts

all: clean docker-run docker-save app