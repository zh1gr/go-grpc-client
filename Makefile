go_mod:
	go mod tidy && go mod vendor && go mod verify

.PHONY: send
send:
	go run client.go