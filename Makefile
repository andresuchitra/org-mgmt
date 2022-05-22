BINARY=engine

unittest: 
	go test -short  ./...

engine:
	go build -o ${BINARY} main.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t org-mgmt .

run:
	docker-compose up --build -d

stop:
	docker-compose down

.PHONY: clean install unittest build docker run stop