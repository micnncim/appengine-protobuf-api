# appengine-protobuf-api

A boilerplate for API with `Google App Engine` and `Protocol Buffers`

## Getting started

```sh
$ find . -type f | xargs sed -i -e 's/micnncim/<REPOSITORY_OWNER>/g' -e 's/appengine-protobuf-api/<REPOSITORY_NAME>/g'
```

Support Protocol Buffers and also JSON.

```sh
$ make run
$ curl -X POST -H 'Content-Type: application/json' localhost:8080/v1/Echo -d '{"body": "hello"}'
{"body":"hello"}
```

## LICENSE

 [MIT](./LICENSE)
