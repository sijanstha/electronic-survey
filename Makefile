build:
	@go build -o bin/evs

uirun:
	@yarn --cwd ./frontend/evs start

uibuild:
	@yarn --cwd ./frontend/evs build

uiinstall:
	@yarn --cwd ./frontend/evs install

uiadd:
	@yarn --cwd ./frontend/evs add $(package)

run: build
	@./bin/evs

test:
	@go test -v ./backend/...

create-migration:
	@migrate create -ext sql -dir db/migration -seq $(name)

container-ls:
	@docker container ls

container-logs:
	@docker container logs $(id)

start:
	@docker compose -f docker-compose.yml up -d --build

stop:
	@docker compose -f docker-compose.yml down

create-network:
	@docker network create evs-network