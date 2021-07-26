# rubberyconf

Free, open, distributed and RESTful configuration engine


# rubberyConf Service

## Run it!

## Prerequisits

Create .env file for docker dependencies like this:
```
    MYSQL_PASSWORD=
    MYSQL_ROOT_PASSWORD=
```
There's a template file in ./docker-env/template.env. Just copy it and rename it as 'local.env'. 

Bootstrap dependencies:
```
     $docker-compose --env-file ./docker-env/local.env up -d
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



