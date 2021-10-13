# Blockchain

## Mod Init

```
go mod init github.com/rafaelcr1dev/Golang-Blockchain.git
```

## Run

```
go run main.go
```

## Run with args

```
go run main.go print

go run main.go add -block "first block"
```

## Run

```
go run main.go createblockchain -address "John"

go run main.go printchain

go run main.go getbalance -address "Lays"

go run main.go send -from "Lays" -to "Fred" -amount 50
```