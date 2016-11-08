run:
	go run app/main.go -p $(postcode)

test:
	go test --race -coverprofile=./tmp/cover.out ./ozdata
	go tool cover -html=./tmp/cover.out -o ./tmp/cover.html

import:
	go run import/main.go
