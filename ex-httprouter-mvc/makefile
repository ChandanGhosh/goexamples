PROJECT=ex-mongo-mvc

build:
	go build -ldflags "-w -s" .

run: build
	./ex-mongo-mvc

docker-build: build
	docker build --rm -t mvc-mongo-goapp .

docker-run: docker-build
	docker run -d -p 8080:8080 mvc-mongo-goapp
