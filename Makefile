clean:
	rm -f docker.sh test
all:
	go build -o docker.sh .