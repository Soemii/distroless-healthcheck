build:
	@echo "Building..."
	@CGO_ENABLED=0 go build -o ./tmp/healthcheck .

clean:
	@echo "Cleaning..."
	@rm -rf ./tmp