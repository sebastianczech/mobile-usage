# Simple app to check mobile usage

## Get IP of gateway

```bash
GATEWAY_IP=`traceroute -w 1 -q 1 -v -m 7 1.2.3.4 2> /dev/null |  awk '$2 ~ /10./ { print $2 }'`
echo $GATEWAY_IP

GATEWAY_PASSWORD="***"
echo $GATEWAY_PASSWORD
```

## Run app 

```
go run main.go $GATEWAY_IP $GATEWAY_PASSWORD
```

## Test app

```
go test
```