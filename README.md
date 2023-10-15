# Go DB Snapshot Replication For MySQL

A *very* basic script to replicate one database and import it into another

## Introduction

All it does, is dump and create a new database.
It doesn't create db credentials or the database server.
Maybe in future it might do more.

## Usage

Copy the `.env.example`

```bash
cp .env.example .env
```

Populate the env to get the original data and create a replica.

Run the bash query.

> You can check the bash script by just running `go run main.go`

```bash
go run main.go | bash -
```

## Resources

As I learn by doing, I'll keep a list of resources that have helped me put this repo together for future reference and as a thank you to those that came before me.

- exec cmd commands. https://zetcode.com/golang/exec-command/
- Dotnev. https://github.com/joho/godotenv
- String interpolation. https://channaly.medium.com/string-interpolation-in-golang-ecb3bcb2d600