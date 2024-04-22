compose-up:
	docker-compose --env-file .env -f docker-compose.yaml up --build --abort-on-container-exit --exit-code-from tax


compose-down:
	docker-compose -f docker-compose.yaml down