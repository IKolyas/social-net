.PHONY: social-network nginx database all down

# Запуск всех сервисов
all:
	docker compose up -d

# Запуск только social-network
social-network:
	docker compose up -d social-network

social-network-down:
	docker compose down social-network

# Запуск только nginx
nginx:
	docker compose up -d nginx

nginx-down:
	docker compose down nginx

# Запуск только database
database:
	docker compose up -d pgmaster pgslave1 pgslave2

database-down:
	docker compose down pgmaster pgslave1 pgslave2

redis:
	docker compose up -d redis

redis-down:
	docker compose down redis

# Остановка всех сервисов
down:
	docker compose down