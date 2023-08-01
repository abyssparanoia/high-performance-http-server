
.PHONY: run-server
run-server:
	go run ./server/main.go ./server/handler.go ./server/context_cancel_middleware.go

.PHONY: run-client
run-client:
	go run ./client/main.go