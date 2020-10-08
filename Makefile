

gosink: main.go
	@go build -o gosink main.go

run:
	@go run main.go

fmt:
	@gofmt -w main.go

install:
	sudo cp gosink /usr/bin

clean:
	@rm -f gosink

