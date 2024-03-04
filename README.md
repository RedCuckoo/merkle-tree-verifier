# merkle-tree-verifier

Client generates files, calculates merkle root and unloads files to the server (deletes locally and uploads to the server).

Server accepts files to store, calculate merkle tree, when requested file, returns merkle proof.

Client when downloading file, can calculate merkle root hash from proof and be sure the file is not corrupted.

---


# Installation

1. Run server:

```shell
docker-compose up -d
```

2. Run client:

```shell
docker run -it redcuckoo/merkle-tree-verifier-client:v0.1.0
```

# Usage

Available commands:

```shell
generate AMOUNT_OF_FILES
list --local
list --remote
unload
download FILE_INDEX
reset
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
