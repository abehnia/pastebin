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