#!/bin/bash

docker-compose --env-file ./docker-env/local.env -f docker-compose.yml up -d
