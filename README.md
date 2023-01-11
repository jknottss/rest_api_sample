# rest_api_sample

notes:
migrate create -ext sql -dir ./schema -seq init   - создание миграций
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up - запуск миграции
