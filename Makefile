.PHONY: client help up down

help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

push: ## Run the server and database
	git push heroku master

open: ## Stop the server and database
	heroku open

logs: ## Start the client
	heroku logs --tail