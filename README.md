# Git Compare [![CI Pipeline](https://github.com/msolimans/gitcomp/actions/workflows/docker-image.yml/badge.svg)](https://github.com/msolimans/gitcomp/actions/workflows/docker-image.yml)

Compare 2 git commits and prints out the difference between those commits

## How to Run

### Using Go
You can build the application if you have `go` installed in your machine. follow the instructions below to build and run

1- Clone the application

```
git clone https://github.com/msolimans/gitcomp
```

2- Build and run
```
go run ./cmd/main.go 
```

### Using Docker

1- Clone the application

```
git clone https://github.com/msolimans/gitcomp
```

2- Build docker image

```
docker build -t gitcomp .
```

3- Run

```
docker run -it --rm gitcomp 
```

### Using Makefile

`Makefile` was added to simplify building and running unit testing in local machine and within CI/CD pipelines

To build locally using installed `go`
```
make REVISION_NUM=sha build
```

To build using `docker`  
```
make REVISION_NUM=sha dbuild
```

To run unit tests
```
make test 
```

## CLI commands and parameters: 

| Command  |  Flag(s)  | Short Flag(s) |  Description             |           Required      |  
|----------|-----------|---------------|------------|-------------|
| help     |  help     |    -h         | Prints help message               |
| version  |           |    -h         | Prints version details               |
| compare  |           |               | Compare 2 git commits               |
|          |  --token  |    -t         | Personal Access Token  | Y |
|          |  --org    |    -o         | Organization or owner | Y |
|          |  --repo   |    -r         | Repostiry name |  Y  |
|          |  --head   |    -d         | Head commit sha  |  Y |
|          |  --base   |    -b         | Base commit sha |  Y |

### Notes

- `structs` were auto-generated using https://mholt.github.io/json-to-go/ and some of them were extended with a `simple` string overrides in `types_ext.go`.
- I used `oath2.NewClient` to generate http client then passed that client inside `GitClient`. we can also pass the token via `Bearer Token` in `Authorization` header as documented here https://docs.github.com/en/rest/commits/commits#compare-two-commits

### Assumptions

- There was no requirement stating anything about pagination but I can follow https://docs.github.com/en/rest/guides/traversing-with-pagination and do that if this is necessary.
- Unit Testing were done based on mocking http test server along with a mock handler that memics the same responses coming from GitHub. GitClient was implemented in away that it can call real or mock Github APIs.
- In case there's a need to change test cases to real GitHub api calls, please let me know and I can adjust that accordingly.
- Http Client returns standard httpClient, enhancement can be done by a simple wrapper within which we can add extra fields like timeouts, unified headers, user agents or logging raw request/response details.
- Nothing mentioned logging that's why I simply used `fmt.Println`. I am pretty familiar with `zap` and `Logrus`
- Requirments for Github actions were not clear enough, I enabled CI for `opened` or `reopened` PRs only however I can add `pushes` to `master` branch however direct pushes can be avoided/prevented from branch protection in repository settings.
### Enhancements

- Add `goreleaser` to package
- Docker build tag can be enhanced with `:date +%s`
- Builds and unit tests can be run in parallel with matrix strategy
