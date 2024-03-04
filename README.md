# merkle-tree-verifier

Client generates files, calculates merkle root and unloads files to the server (deletes locally and uploads to the server).

Server accepts files to store, calculate merkle tree, when requested file, returns merkle proof.

Client when downloading file, can calculate merkle root hash from proof and be sure the file is not corrupted.

---

# Requirements

* [Docker ^20.10.6](https://www.docker.com/get-started)
* [Compose ^3.3](https://docs.docker.com/compose/install/)
* [Go ^1.20](https://golang.org/)


# Installation

1. Run server:

```shell
docker-compose up -d
```

2. Run client:

```shell
docker pull redcuckoo/merkle-tree-verifier-server:v0.2.0
docker run --network merkle-tree-verifier_default -it redcuckoo/merkle-tree-verifier-client:v0.2.0
```

## OR

Run script:
```shell
./run.sh
```

# Usage

Available commands:

```shell
generate AMOUNT_OF_FILES # generate N amount of files on the client
list --local # list files stored locally on the client
list --remote # list files stored remotely on the server
unload # send files to the server, store merkle root, delete local
download FILE_INDEX # download file from the server
reset # reset client and server
exit
```

Simple demo:
```shell
generate 10
list --local
list --remote
unload
list --local
list --remote
download 1
list --local
list --remote
reset
```
