#!/bin/bash

docker-compose up -d

docker pull redcuckoo/merkle-tree-verifier-server:v0.2.0

docker run --network merkle-tree-verifier_default -it redcuckoo/merkle-tree-verifier-client:v0.2.0

docker-compose down -v