postgres:
	docker run --name postgreslatest -p 5432:5432 -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=staff -d postgres:latest    

postgresstop:
	docker stop postgreslatest

postgresrm:
	docker rm postgreslatest

createdb:
	docker exec -it postgreslatest createdb --username=staff --owner=staff simplebank

dropdb:
	docker exec -it postgreslatest dropdb --username=staff simplebank

migrateup:
	migrate -path db/migration -database "postgresql://staff:secret@localhost:5432/simplebank?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://staff:secret@localhost:5432/simplebank?sslmode=disable" -verbose down 

sqlc:
	sqlc generate 

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createdb dropdb postgres postgresstop postgresrm migrateup migratedown sqlc server

