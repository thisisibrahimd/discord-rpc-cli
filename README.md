# Discord-RPC-CLI

Discord-RPC-CLI is a cross-platform cli tool to set discord rich presences.

## Install

1. Clone repo

```
$ go get github.com/thisisibrahimd/discord-rpc-cli
```
2. Create executable
```bash
$ cd discord-rpc-cli
$ go install drpc.go
```

## Usage

Discord must be open before running executable

Example:
```bash
$ drpc set 801624440328552458 --lgimg dragon --lgtxt "Using DragonAIO" --details "Using DragonAIO" --state "Version 9.7.0"
```

## License
[MIT](https://choosealicense.com/licenses/mit/)
