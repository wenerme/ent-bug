gen:
	go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema

gen-test: gen test

.PHONY: test
test:
	cd test && go test -v -run TestClear

test-pg:
	docker rm -f pg-test
	docker run -d \
      -e POSTGRES_USER=test \
      -e POSTGRES_PASSWORD=test \
      -e POSTGRES_DB=test \
      --name pg-test -p 15432:5432 postgres:alpine
	cd test && go test -v -run TestPG
	docker rm -f pg-test
