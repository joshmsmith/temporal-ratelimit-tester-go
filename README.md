# Rate Limit Tester
Sample golang app with workers and clients to generate load.


## Temporal Cloud Connection
This example assumes that you have a temporal cloud namespace configured and have local client certificate files for your namespace.
It also assumes the usage of [direnv](https://direnv.net/).

The values are passed into the demo app using environment variables, example direnv .envrc file is included in the repo:

```
# direnv .envrc

# Temporal Cloud connection
# region: us-east-1
export TEMPORAL_HOST_URL="myns.abcdf.tmprl.cloud:7233"
export TEMPORAL_NAMESPACE="myns.abcdf"

# tclient-myns client cert
export TEMPORAL_TLS_CERT="/Users/myuser/.temporal/tclient-myns.pem"
export TEMPORAL_TLS_KEY="/Users/myuser/.temporal/tclient-myns.key"

# Optional: path to root server CA cert
export TEMPORAL_SERVER_ROOT_CA_CERT=
# Optional: Server name to use for verifying the server's certificate
export TEMPORAL_SERVER_NAME=

export TEMPORAL_INSECURE_SKIP_VERIFY=false

# App temporal taskqueue name for moneytransfer
export TRANSFER_MONEY_TASK_QUEUE="go-moneytransfer"
# timer for transfer table to be checked (seconds)
export CHECK_TRANSFER_TASKQUEUE_TIMER=20

# payload data encryption
export ENCRYPT_PAYLOAD=false
export DATACONVERTER_ENCRYPTION_KEY_ID=mysecretkey

# Set to enable debug logger logging
export LOG_LEVEL=debug

# local mysql backend db connection
export MYSQL_HOST=localhost
export MYSQL_DATABASE=dataentry
export MYSQL_USER=mysqluser
export MYSQL_PASSWORD=mysqlpw
```

## Sample usage
It was valuable to me to decrease the limits of my namespace below the defaults, to generate rate-limiting throughput without overloading my testing machine (i7-13700H   2.40 GHz/32MB RAM) or incur unnecessary spend on infrastructure. 
I went with 40 APS/160 RPS. 
1. Start workers 

```
go run accumulator/worker/main.go
```
I recommend starting several, I went with 6.

2. Start workflows
```
go run accumulator/starter/main.go
```
I recommend starting several, I ramped up from 2 up to 7 to get above 160 RPS.

3. Start Querier
```
go run accumulator/query/main.go
```
I recommend 2.

4. Monitor your namespace's metrics, observe rate limiting.