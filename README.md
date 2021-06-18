# rubberyconf

Free, open, distributed and RESTful configuration engine


# rubberyConf Service

## Prerequisits

Create .env file for docker dependencies like this:
```
    MYSQL_PASSWORD=
    MYSQL_ROOT_PASSWORD=
    MONGO_INITDB_ROOT_PASSWORD=
```
Bootstrap dependencies:
```
     $docker-compose up -d --env-file ./docker.env
```
## rubberyConf Api

### Configuration syntax

Examples of features:
 * docs/toggle1.yml - simple feature
 * docs/toggle2.yml - advanced example
 * docs/toggle3.yml - other example 

### Set up your api

* /config/local.yml

### Endpoint

- GET /conf/:feature 
- GET /conf/:feature/:branch

Get value stored in :feature



