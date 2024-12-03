.PHONY: run-app build-all stop-app rerun-app get-swagger


run-app: build-all
	docker compose up -d ls-app ls-postgres

build-all:
	docker compose build

stop-app:
	docker stop ls-app ls-postgres

rerun-app:
	make stop-app && make run-app

get-swagger:
	swag init -g cmd/app/main.go -o ./docs_swagger
