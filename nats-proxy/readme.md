go run nats-proxy/proxy.go

go run nats-proxy/client.go

curl http://127.0.0.1:8080/user/info
{"Name":"Alan"}