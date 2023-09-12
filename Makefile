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
