# Super Pastebin

It's pastebin, but super.

Install postgres 16, make, then:

```
make stop-db start-db db-shell
```

then

```
export DB_STRING="host=/home/ps/src/telegraf-http-sink user=telegraf-http-sink_admin dbname=telegraf-http-sink"
go run .
```

## How to run the server

```
go run main.go
```
## How to test the functionality

Create a new bin

```
curl -X POST http://localhost:8080/bins -H "Content-Type: application/json" -d '{"text": "that's my bin y'all"}'
```

Get an existing bin
```
curl http://localhost:8080/bins/put-the-uuid-of-the-bin-here
```
