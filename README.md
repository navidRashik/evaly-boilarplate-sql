# Go Boilerplate

## Start Rest 

## Build
```bash
$ ./build.sh
or
$ make build
```


## Application binary (run server)
```bash
$ go-boilerplate serve -c example.config.yaml
or
$ make run
```

## Container dev
```bash
$ docker-compose up --build
or
$ make serve
```

## GuideLine

* api folder contains rest code
* rpcs folder contains grpc code
* rpcrestproxy contains grpc rest proxy code

* infra contains drivers like db, messaging, cache etc
* repo folder contains database code
* model folder contains model
* service folder contains application service

### flow
> cmd -> api/rpcs/rpcrestproxy -> service -> repo, models, cache, messaging


### Example APIS

### Health 

Method: `GET`
URL: `http://{base_url}:{system_server_port}/system/v1/health/api`


Response:
```status_code: 200```
```json=
{
"data": "ok"
}
```