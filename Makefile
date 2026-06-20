all:
	go build -o . ./...

clean:
	rm <your_application>

.PHONY: all clean
