# Travels Budget

Tool to build a personalized itinerary planner based on **budget and number of days** to inspire travelers to seek out the memorable things to do.

## Table of Contents

- [Getting started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Application init](#application-init)
    - [Development](#development)
    - [Testing](#testing)
- [Solution Space](#solution-space)
    - [Demo](#demo)
    
## Getting started

### Prerequisites

You need to make sure that you have `make`, `docker` and `go` installed

```
$ which make
/usr/bin/make
$ which docker
/usr/local/bin/docker
$ which go
/usr/local/bin/go
```

There is no other configuration needed in order to setup this project for development.

### Application init

The first thing you need to do is init the application, create the docker image. To do so, run the command
```
make docker-init
```

It will create the image `travels-budget`.

### Development

Routine operations are defined in `Makefile`.

```
 Base commands to work/dev the tool to inspire travelers to seek out the memorable things to do.

 Working with local Go:
  deps                                          - Ensures dependencies using dep and installs several required tools
  build                                         - Build the application
  test                                          - Run all tests

 Working with Docker:
  docker-init                                   - Build docker image with name travels-budget
  docker-test                                   - Run tests using docker

 Running algorithm:
  itinerary-planner                             - Runs the algorithm using the docker image travels-budget

 Arguments

  ACTIVITIES                                    JSON file with a subset of activities used to suggest the itinerary.
  BUDGET                                        Budget willing to spend in the activities
  DAYS                                          Days willing to spend in the activities

 Usage

  make itinerary-planner ACTIVITIES=<filepath> BUDGET=<budget> DAYS=<days>

 Example

  Runs the algorithm using the values define in the problem using the docker image travels-budget
  alias "travels-budget personalized -b 680 -d 2" using the file berlin.json

  make demo
```

[[table of contents]](#table-of-contents)

### Testing

We highly recommend every time you want to run the test suite, update your dependencies `make deps`.

```
make test
```

Using docker (update the dependencies inside the container)

```
make docker-test
```

## Solution Space

To build a personalized itinerary planner based on budget and number of days run

```
make itinerary-planner ACTIVITIES=<filepath> BUDGET=<budget> DAYS=<days>
```

`BUDGET`: 
   * integer between 100 and 2000. 
   * The average budget to be allocated per day needs to be 50. Which means for budget 100 you can only have 1 or 2 days. Which also means the minimum budget for 5 days needs to be 250.

`DAYS`: 
   * integer between 1 and 5.

`ACTIVITIES`: 
   * string file path activities given by the product team


**Note** We highly recommend every time before you used the command for the first time, to run `docker-init` to re-build the docker image with the latest code.

### Demo

To check the tool working out of the box run (it will display random itineraries each team)

```
make demo
```

`BUDGET`: 680 

`DAYS`: 2 

`ACTIVITIES` JSON file with a subset of activities from Berlin (1000 in total) http://tiny.cc/gyg-berlin-01

[[table of contents]](#table-of-contents)
