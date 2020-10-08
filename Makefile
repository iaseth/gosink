

gosink: main.go
	@go build -o gosink main.go

install:
	sudo cp gosink /usr/bin

clean:
	@rm -f gosink

