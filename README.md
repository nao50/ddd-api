# usage

```sh
$ go run cmd/main.go
$ curl -X POST -H "Content-Type: application/json" -d '{"name": "nao", "tag": "fish"}' localhost:5050/pets
$ curl localhost:5050/pets
$ curl localhost:5050/pets/{id
$ curl -X PUT -H "Content-Type: application/json" -d '{"name": "nanao", "tag": "tree"}' localhost:5050/pets/{id}
$ curl -X DELETE localhost:5050/pets/{id}
```
