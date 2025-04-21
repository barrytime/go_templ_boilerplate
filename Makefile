.PHONY: clean
clean:
	rm -rf tmp/
	rm -f ./static/css/main.css
	rm -f ./static/js/*.js

db_login:
	psql -U admin -d dev_db -h localhost -p 5432

.PHONY: db_up db_down db_reset

db_up:
	go run ./cmd/migrate/main.go up

db_down:
	go run ./cmd/migrate/main.go down


db_reset: db_down db_up

.PHONY: templ
templ:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate -watch

.PHONY: css-watch
css-watch:
	npx @tailwindcss/cli -i ./frontend/css/input.css -o ./public/css/main.css --watch

.PHONY: js-watch
js-watch:
	bun build --outdir ./public/js --target node ./frontend/js/*.js --watch

.PHONY: dev
dev:
	@air & templ generate -watch