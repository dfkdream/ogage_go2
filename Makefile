all:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o . ./...

run:
	go run cmd/ogage_go2/main.go

clean:
	rm ogage_go2

.PHONY: all run clean
