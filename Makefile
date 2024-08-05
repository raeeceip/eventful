.PHONY: all backend frontend clean build

all: backend frontend

backend:
	@echo "Building backend..."
	cd eventful-backend && go build -o ../eventful-electron/resources/backend

frontend:
	@echo "Building frontend..."
	cd eventful-electron && npm install
	cd eventful-electron && npm run build

clean:
	@echo "Cleaning up..."
	rm -f eventful-electron/resources/backend
	rm -rf eventful-electron/dist
	rm -rf eventful-electron/node_modules

build: clean all
	@echo "Build complete!"