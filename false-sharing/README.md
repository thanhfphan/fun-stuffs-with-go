# False sharing pattern

## Naive

```bash
cd ./naive
go build .
time ./naive
```

## False sharing

```bash
cd ./false-sharing
go build .
time ./false-sharing
```

## Fix false sharing(magic happens here :D)

```bash
cd ./fix-false-sharing
go build .
time ./fix-false-sharing
```
