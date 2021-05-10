build:
	docker-compose -f docker-compose.yml -f docker-compose.kafka.yml build

up:
	docker-compose -f docker-compose.yml -f docker-compose.kafka.yml up -d

ps:
	docker-compose -f docker-compose.yml -f docker-compose.kafka.yml ps

