.PHONY: dependency unit-test integration-test docker-up full-test docker-down clear

dependency:
	@go get -v ./...

integration-test: dependency
	@go test -v ./... -p 1 -cover -coverprofile=coverage.out

unit-test: dependency
	@go test -v -short ./... -p 1

report-test:
	@go tool cover -html=coverage.out

full-test:
	@docker-compose -f docker-compose-test.yml up --build --abort-on-container-exit

docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down

clear: docker-down