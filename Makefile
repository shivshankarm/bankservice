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

migrateup1:
	migrate -path db/migration -database "postgresql://staff:secret@localhost:5432/simplebank?sslmode=disable" -verbose up 1 

migratedown:
	migrate -path db/migration -database "postgresql://staff:secret@localhost:5432/simplebank?sslmode=disable" -verbose down 

migratedown1:
	migrate -path db/migration -database "postgresql://staff:secret@localhost:5432/simplebank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate 

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/shivshankarm/bankservice Store	

.PHONY: createdb dropdb postgres postgresstop postgresrm migrateup migratedown sqlc server migrateup1 migratedown1 mock

