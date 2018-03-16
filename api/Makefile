# sudo apt install build-essentials

SSL_PATH=$(shell pwd)/server/ssl

help:
	@echo "docker-api     Run Go Api"
	@echo "docker-build   Build Go image"
	@echo "docker-mongo   Run Mongo DB server (mongodb://mongodb:27017)"
	@echo "format         Format Go packages"
	@echo "gencert        Generate SSL certificate"

gencert:
	@openssl genrsa -out $(SSL_PATH)/server.key 2048
	@openssl req \
        -new \
        -x509 \
        -sha256 \
        -days 365 \
        -subj "/C=FR/ST=IDF/L=Paris/O=ipssi/CN=localhost" \
        -key $(SSL_PATH)/server.key \
        -out $(SSL_PATH)/server.cert

format:
	@go fmt $(shell go list ./...)

docker-mongo:
	@docker run --name mongodb -p 127.0.0.1:27017:27017 -d mongo

docker-build:
	@docker rmi -f ipssi/api:1.0 
	@docker build -t ipssi/api:1.0 $(shell pwd)

docker-api:
	@docker run --name api -p 3000:3000 --link mongodb:mongodb -d ipssi/api:1.0