.PHONY: run
run:
	API_PORT=8080 \
	DATABASE_URI=postgres://postgres:changeme@localhost:5432/pmd \
	LOG_LEVEL=DEBUG \
	go run ./cmd/web-api-reference
