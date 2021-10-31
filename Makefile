build-backend-image:
	docker build -t api -f ./infra/Dockerfile.backend .

run-backend-server:
	docker run --rm -it -p 8888:8080 api