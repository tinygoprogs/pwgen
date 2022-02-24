# pwgen
generates passwords, complies with nasty rules (e.g. password must contain one
uppercase A and no lowercase n)

## usage
```sh
go install github.com/tinygoprogs/pwgen/cmd/pwgen
```

Alternative:
```sh
go get -u github.com/tinygoprogs/pwgen
cd $GOPATH/src/github.com/tinygoprogs/pwgen
go run cmd/pwgen
```

## RNG
uniformity tested via
```sh
go build cmd/main.go
while true; do ./main -len 64 2>>data; done
^C
python -c 'from collections import Counter as C; fd=open("data"); print(C(fd.read()).most_common(1000))'
```
