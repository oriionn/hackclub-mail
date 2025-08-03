build:
	go build -o hackclub-mail main.go help.go fetch.go utils.go

test:build 
	./hackclub-mail