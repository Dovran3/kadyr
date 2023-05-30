run:
	go run main.go
db:
	psql -U postgres -d postgres -f ./database/init.sql
	psql -U postgres -d kadyr -f ./database/create.sql