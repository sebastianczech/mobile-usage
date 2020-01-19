# Simple app to check mobile usage

## File with credentials

```bash
export GATEWAY_IP=`traceroute -w 1 -q 1 -v -m 7 1.2.3.4 2> /dev/null |  awk '$2 ~ /10./ { print $2 }'`
echo $GATEWAY_IP

export GATEWAY_PASSWORD="***"
echo $GATEWAY_PASSWORD

export NJU_LOGIN="***"
echo $NJU_LOGIN

export NJU_PASSWORD="***"
echo $NJU_PASSWORD
```

## Run app 

```
. credentials.sh 
go run main.go $GATEWAY_IP $GATEWAY_PASSWORD $NJU_LOGIN $NJU_PASSWORD
```

## Test app

```
go test
```