# Variables
IMAGE_NAME = go-sqlite-app

build:
	docker build -t $(IMAGE_NAME) .

run:
	docker run --rm -v ${PWD}/data:/app/data $(IMAGE_NAME)

clean:
	docker rmi $(IMAGE_NAME)
