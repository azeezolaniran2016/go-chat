dev-start:
	docker-compose up -d --build
	
dev-init-deps:
	govendor init

dev-logs:
	docker-compose logs -f go-chat

dev-clean:
	docker-compose down --rmi all

dev-stop:
	docker-compose stop

dev-fetch-deps:
	@echo "Fetching Dependencies..."
	govendor fetch github.com/gorilla/mux
	govendor fetch github.com/sirupsen/logrus
	@echo "Dependencies fetched!"
