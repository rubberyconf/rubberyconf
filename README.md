# rubberyconf

Free, open, distributed and RESTful configuration engine

`rubberyconf` provides an HTTP service to enable business features on-demand. It is useful when we want to use feature toggles to speed up the development cycle or change business logic without recompile/release new code.

# rubberyConf Service

## Run it!

- First of all start dependencies
```
docker-compose up -d
```
If you need pass sensible information you can use *./dockerup.sh* script.

- Second create your own configuration file in `./conf` folder. For example: 
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

- Lastly, run it: 

``` 
$cd ./cmd/server
$ENV=local go run main.go
```

You'll be able to read a message in console that `rubberyconf`is running in port 8080.

### Create your first configuration

Now, when create our own configuration:

```
curl -X POST "http://localhost:8080/conf/myfirstfeature"
{
    "default": {
        "value": {
            "data": "hello world",
            "type": "string"
        }
    }
}

```
Every configuration has a default value. Where value can be an "string", "number" or "json". In *data* field is the value and *type* field is vaule's type (meta information). 

You can get your feature: 

```
curl -X GET http://localhost:8080/conf/myfirstfeature
```
And you'll get this result with http code 200:
```
hello world
```

## rubberyConf Api

### Endpoints

- GET /feature/:feature 
- GET /conf/:feature 
- PATCH /conf/:feature 
- POST /conf/:feature
- DELETE /conf/:feature

[Swagger file description](./docs/api/swagger.yml)

### Feature description 

In folder *./docs/features* you'll find some examples of features addmited by the rubberyconf.

Examples of features:
 * [toggle1.yml - simple feature](./docs/features/toggle1.yml)
 * [toggle2.yml - advanced example](./docs/features/toggle2.yml)
 * [toggle3.yml - other example](./docs/features/toggle3.yml) 


## Datasource

`rubberyconf` allows persist your configurations at: 

- InMemory (just for develop and test)
- Mongo
- Gogs
- GitHub

## Cache

To enable this configurations extremly fast for your consumers `rubberconf` uses a cache system.

- None (cache is disabled)
- Redis
- InMemory (in-process memory)

# Configuration-as-code

`rubberyconf` connect your clients (web, apps or other services) with your configuration. It allows changing the configuration dynamically. For example, it is useful when you're building a **feature toggle** system. You release your changes behind a feature toggle through your CI/CD. Releasing as many times as you need, but having the feature toggle (aka configuration) disabled. Then, when you want to start releasing or testing your changes, you can change the configuration for a concrete set of users. Initially, all clients will receive the default value, i.e., **off**. However, a subset of your clients can receive another value like **on**. So, you can start testing with a subset of real clients and progressively roll it out by country, locations, flags and so on.


## rubberyconf with Github.com and Redis

This configuration allows you to store your configurations in a Githug repo. `rubberyconf` will reach Github to retrieve the configuration and it will store it in *redis* for some TTL. It reduces the impact on Github service based on TTL value.  The main advantages of this configuration are: you can use *pull request* and *commits* to track changes in your configuration with your favourites tools. 

Instead of using Github, you can use [Gogs](https://github.com/gogs/gogs).


# Logics applied to configurations

Every configuration has a set of logics, those logics can change default value if some circumstances apply to each client request. 

Examples: 
- Environment (local, stage, production, ...)
- http headers (ip, browser, platform, location, etc..)
- http querystring parameter
- client version

A complete list of features implemented is [here](./internal/feature/rules)

Those logics are extensibles, your contributions are welcome.   

# License

[Apache 2.0 license](./LICENSE)

