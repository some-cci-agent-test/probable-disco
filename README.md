# probable-disco 

Minimal Go module with a flaky test and CircleCI setup.

## Local Development

- Initialize and tidy modules (already done):

```bash
go mod tidy
```

- Run tests with flakiness disabled for stability:

```bash
FLAKY_OFF=1 go test -v ./...
```

- Run tests with flakiness enabled (may fail ~30% of runs):

```bash
go test -v ./...
```

## CI

CircleCI config at `.circleci/config.yml` runs `go test -v ./...` on each push.
