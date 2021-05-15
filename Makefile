build:
	docker-compose -f docker-compose.yml -f docker-compose.kafka.yml build

up:
	docker-compose -f docker-compose.yml -f docker-compose.kafka.yml up -d

ps:
	docker-compose -f docker-compose.yml -f docker-compose.kafka.yml ps

build-debug:
	docker-compose -f docker-compose.debug.yml -f docker-compose.kafka.yml build

up-debug:
	docker-compose -f docker-compose.debug.yml -f docker-compose.kafka.yml up -d

ps-debug:
	docker-compose -f docker-compose.debug.yml -f docker-compose.kafka.yml ps


