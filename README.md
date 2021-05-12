# logger

> Drewis Dorner

Simple go application to use inside a container, which accepts all http requests and outputs the data of the `RequestBody` to the console.

## Build

```bash
docker build . -t logger
```

## Use

Send HTTP Request to Container

```bash
curl -X POST -H "Content-Type: application/json" \
    -d '{"name": "tester", "email": "supertester@example.com"}' \
    localhost:8080
```

Receive the following output from the container

```
1970/01/01 00:00:00 / : {"name": "tester", "email": "supertester@example.com"}
```