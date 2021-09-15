# BetmeAPI
API written for Betme base on Odds API V3.
Store data to a mongodb which needs user to indicate in a commandline.
Fetch data can be used to feaure build REST or grpc apis.

## Usage

```bash
go build main.go
main -k [APIKEYHERE] -a [MONGO ADDRESS HERE]
```

## License
