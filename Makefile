all:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -trimpath -ldflags '-s -w -extldflags "-static"' -o . ./...

run:
	go run cmd/ogage/main.go

release: all
	mkdir ogage_go2
	cd ogage_go2; \
	cp ../ogage .; \
	cp ../scripts/*.sh .;
	zip ogage_go2.zip ogage_go2/*
	rm -r ogage_go2

clean:
	rm ogage evtest ogage_go2.zip

.PHONY: all run release clean
