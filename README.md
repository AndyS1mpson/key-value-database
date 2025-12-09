# key-value-database

## Introduction

In-memory key-value database with support of asynchronous replication (physic) and WAL (Write-Ahead Logging)

## Launch database server

### Project configuration

You need to create `.config.yaml` file and fill all variables like in `.config.yaml.example`


### Launch
```
make run-server
```

## Launch database cli

You can launch database cli:
```
make run-cli 
```

with follow command line arguments:
- address           - Address of the database
- idle_timeout      - Idle timeout for connection
- max_message_size  - Max message size for connection
