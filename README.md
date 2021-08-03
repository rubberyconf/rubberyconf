# rubberyconf

Free, open, distributed and RESTful configuration engine

`rubberyconf` provides an HTTP service to enable business features on-demand. It is useful when we want to use feature toggles to speed up the development cycle or change business logic without recompile/release new code.


# rubberyConf Service

## Run it!

First of all start dependencies
```
docker-compose up -d
```
Second create your own configuration file in `./conf` folder. For example: 
```
api:
  port: 8080
  cache: "None" 
  source: "InMemory"
  logs: 
    - Console
  default_ttl: "3600s"
  options:
    loglevel: "debug"  
```
If you're running in your local environment you must name this file: **local.yml**. 

Last, run it: 

``` 
$cd ./cmd/server
$ENV=local go run main.go
```

You'll be able to read a message in console that `rubberyconf`is running in port 8080.

### Create your first configuration

```
curl -X POST "http://localhost:8080/conf/myfirstfeature"
{
    "name": "feature2",
    "default": {
        "value": {
            "data": "hello world",
            "type": "string"
        }
    }
}

```
Every configuration has a default value. Where value can be an "string", "number" or "json". In *data* field is the value and *type* is vaule's type. 

Now, you can get the of your feature

```
curl -X GET http://localhost:8080/conf/myfirstfeature
```
And you'll get this result with http code 200:
```
hello world
```

## rubberyConf Api

### Configuration syntax

In folder *./docs_examples* you'll find some examples of features addmited by the rubberyconf.

[Swagger file description](./docs_examples/swagger.yml)

Examples of features:
 * toggle1.yml - simple feature
 * toggle2.yml - advanced example
 * toggle3.yml - other example 
 * toggle4.yml - another example 


### Endpoints

- GET /conf/:feature 
- GET /conf/:feature/:branch (if you're using Gogs or Github)
- POST /conf/:feature
- DELETE /conf/:feature

## Datasource

`rubberyconf` allows read/store your configurations at: 

- InMemory (just for develop and test)
- Mongo
- Gogs
- GitHub

## Cache

To make this configurations extremly fast `rubberconf` uses a cache storage.
- None
- Redis
- InMemory (in-process memory)






