NODE_BIN := ./node_modules/.bin

run:build
	@./.bin/epiqpixai

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss
	@npm install -D daisyui@latest



build:
	$(NODE_BIN)/tailwindcss -i view/css/app.css -o public/styles.css
	@templ generate view
	@go build -o ./.bin/epiqpixai main.go


up: ## DB migration up
	@go run cmd/migrate/main.go up

drop:
	@go run cmd/drop/main.go up

down: ## DB migration down
	@go run cmd/migrate/main.go down

migration : ## migrations against the DB
	@migrate create -ext sql -dir cmd/migrate/migrations ${filter-out $@, ${MAKECMDGOALS}}

seed:
	@go run cmd/seed/main.go