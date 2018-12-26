# This makefile contains some convenience commands for deploying and publishing.

# Build docker images for both server and client apps
all: srv cli

# Build docker image with server app
srv:
	cd server && \
	docker build -t ivchms:server .

# Build docker image with client app
cli:
	cd client && \
	docker build -t ivchms:client .

# Remove docker images
clean:
	docker rmi ivchms:server -f
	docker rmi ivchms:client -f