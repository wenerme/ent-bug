gen:
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature graphql ./ent/schema

gen-test: gen test

.PHONY: test
test:
	cd test && go test -v -run TestClear
test-pg:
	cd test && go test -v -run TestPG
gen-test-pg: gen test-pg



test-pg-unit:
	docker rm -f pg-test
	docker run -d \
      -e POSTGRES_USER=test \
      -e POSTGRES_PASSWORD=test \
      -e POSTGRES_DB=test \
      --name pg-test -p 15432:5432 postgres:alpine
	cd test && go test -v -run TestPG
	docker rm -f pg-test

start-pg:
	docker run -d \
      -e POSTGRES_USER=test \
      -e POSTGRES_PASSWORD=test \
      -e POSTGRES_DB=test \
      --name pg-test -p 15432:5432 postgres:alpine

stop-pg:
	docker rm -f pg-test

tidy:
	go mod tidy