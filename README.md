# go-grpc-client

---

Golang example project represent gRPC client.

[Golang Server](https://github.com/zh1gr/go-grpc-server)<br>
[Golang Domain](https://github.com/zh1gr/go-grpc-domain)

### install

---

```shell
go run client.go
```

### commands

---

`make go_mod` - update dependencies and store it in **vendor** directory in project

`make send` - Run client to proceed test message to Server.<br>


### project files

---

 - **server.go** - Contains the implementation of the client-side logic for interacting with the gRPC service.
 - **Makefile** - Script containing commands for automating common development tasks such as building, testing, and running the project.
 - **go.mod** - File defining the module and dependencies for the Go project.
 - *vendor* - Directory containing third-party dependencies managed by a dependency management tool like Go modules or dep.