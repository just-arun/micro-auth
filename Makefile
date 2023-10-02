dev:
	go run ./main.go serve -e dev -p 8090 -c .
access:
	go run ./main.go access -e dev -c .
migrate:
	go run ./main.go migrate -e dev -c .
env:
	go run ./main.go env
keys:
	go run ./main.go keys
setup: 
	make migrate && make access

