run: gwire build

gwire: 
	wire gen ./wire

build:
	go run main.go