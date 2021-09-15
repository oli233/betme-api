# BetmeAPI
API written for Betme base on Odds API V3.
Store data to a mongodb which needs user to indicate in a commandline.
Fetch data method can be used to feaure build REST or grpc apis.

Two apis:
```go
type Api interface {
	InitData() int
	FetchData() int
}

```
## Usage

```bash
go build main.go
main -k [APIKEYHERE] -a [MONGO ADDRESS HERE]
```

## License
