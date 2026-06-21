all:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -trimpath -ldflags '-s -w -extldflags "-static"' -o . ./...

run:
	go run cmd/ogage/main.go

clean:
	rm ogage evtest

.PHONY: all run clean
