

gosink: main.go
	@go build -o gosink main.go

install:
	@cp gosink /usr/bin

clean:
	@rm gosink

