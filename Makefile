build:
	./build.sh

run: build
	go run main.go serve --config ./example.config.yaml

serve:
	docker-compose down
	docker-compose up -d
