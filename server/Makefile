run:
	go build -o browser-remote-server
	./browser-remote-server -config config/local.yaml

swag:
	mkdir -p ./docs/api/
	swagger generate spec -o ./docs/api/swagger.json
	swagger serve ./docs/api/swagger.json