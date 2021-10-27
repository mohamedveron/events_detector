# events_detector

## Assumtions:
I store the event for each struct in memory cached map to follow its construction until completion

## Setup of the component

Must have golang installed version >= 12.0.0

make file consists of 4 steps: generate, test, build, run
you can run all of them 

```bash
make all
```



## Test consume accounts data and make transferes by run

```bash
make test
```
or

go test -v ./tests


# Start the http server on port 9090:

```bash
make run
```

# Build and run docker image

```bash
docker build --tag events_detector .
```

```bash
docker run -p 9090:9090 events_detector
```
