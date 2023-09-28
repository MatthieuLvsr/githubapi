restart:
	@echo "Restarting the server..."
	@docker-compose down
	@docker rmi -f githubapi-app
	@docker-compose up -d --build

log:
	@echo "Showing logs..."
	@docker logs -f githubapi-app-1