run:
	go run main.go
db:
	psql -U postgres -d postgres -f ./database/init.sql
	psql -U postgres -d ave_db -f ./database/create.sql
