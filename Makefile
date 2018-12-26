all: srv cli

srv:
	cd server && \
	docker build -t ivchms:server .

cli:
	cd client && \
	docker build -t ivchms:client .

clean:
	docker rmi ivchms:server
	docker rmi ivchms:client