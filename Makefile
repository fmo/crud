create-migration:
	migrate create -ext sql -dir migrations -seq create_tasks_table

migrate:
	migrate -database "postgres://user:password@localhost:5433/tasks?sslmode=disable" -path ./migrations up

migrate-down:
	migrate -database "postgres://user:password@localhost:5433/tasks?sslmode=disable" -path ./migrations down

migrate-force:
	migrate -database "postgres://user:password@localhost:5433/tasks?sslmode=disable" -path ./migrations force 1
