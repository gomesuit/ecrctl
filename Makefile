PROJECT_NAME=ecrctl
VERSION=v0.1

all: clean bin

bin:
	go build -o bin/$(PROJECT_NAME)
	
clean:
	rm -rf bin

