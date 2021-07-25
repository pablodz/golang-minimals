

1. go run main.go
2. Execute command
   
```bash
curl -X POST \
	http://localhost:8000/v1/user \
	-H 'cache-control: no-cache' \
	-H 'content-type: application/json' \
	-d '{
	"username":"pedro",
	"email_address":"pedro@example.com",
	"first_name":"Pedro",
	"last_name":"Martinez"
	}'
```