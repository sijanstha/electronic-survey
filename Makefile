build:
	@go build -o bin/evs

uirun:
	@yarn --cwd ./frontend/evs start

uibuild:
	@yarn --cwd ./frontend/evs build

uiinstall:
	@yarn --cwd ./frontend/evs install

run: build
	@./bin/evs

test:
	@go test -v ./backend/...

container-ls:
	docker container ls

container-logs:
	docker container logs $(id)

start:
	docker compose -f docker-compose.yml up -d --build

stop:
	docker compose -f docker-compose.yml down