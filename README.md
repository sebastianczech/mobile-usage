# Simple app to check mobile usage

## Get IP of gateway

```bash
GATEWAY=`traceroute -w 1 -q 1 -v -m 7 1.2.3.4 2> /dev/null |  awk '$2 ~ /10./ { print $2 }'`
echo $GATEWAY
```

## Run app 

```
go run main.go 
```

## Test app

```
go test
```