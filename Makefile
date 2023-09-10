dev:
	go run ./main.go serve -e dev -p 8090 -c .
migrate:
	go run ./main.go migrate -e dev -c .