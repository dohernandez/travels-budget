# --------------------------------------------------------------------------------------------------------------------
# Variables
# (https://www.gnu.org/software/make/manual/html_node/Using-Variables.html#Using-Variables)
# --------------------------------------------------------------------------------------------------------------------
branch = dev
VERSION ?= dev
revision = dev
build_user = $(USER)
build_date = $(shell date +%FT%T%Z)

# detecting GOPATH and removing trailing "/" if any
GOPATH = $(realpath $(shell go env GOPATH))
IMPORT_PATH = $(subst $(GOPATH)/src/,,$(realpath $(shell pwd)))
export APP_NAME ?= $(subst github.com/dohernandez/,,$(IMPORT_PATH))

VERSION_PKG = $(IMPORT_PATH)/vendor/github.com/dohernandez/$(APP_NAME)/pkg/version
export LDFLAGS = -X $(VERSION_PKG).version=$(VERSION) -X $(VERSION_PKG).revision=$(revision) -X $(VERSION_PKG).buildUser=$(build_user) -X $(VERSION_PKG).buildDate=$(build_date)

BUILD_DIR ?= bin/
BINARY=cli

# Change this variables

# Colorz
NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

# Filters variables
CFLAGS=-g
export CFLAGS

all: help

## -- Working with local Go --

run: clean deps build

#-----------------------------------------------------------------------------------------------------------------------
# House keeping - Cleans our project: deletes binaries
#-----------------------------------------------------------------------------------------------------------------------
clean:
	@printf "$(OK_COLOR)==> Cleaning build artifacts$(NO_COLOR)\n"
	@rm -rf $(BUILD_DIR)

# --------------------------------------------------------------------------------------------------------------------
# Dependencies
# --------------------------------------------------------------------------------------------------------------------
## Ensures dependencies
deps:
	@printf "$(OK_COLOR)==> Installing dep$(NO_COLOR)\n"
	@test -s $(GOPATH)/bin/dep || go get -u github.com/golang/dep/cmd/dep
	@$(GOPATH)/bin/dep ensure -v

# --------------------------------------------------------------------------------------------------------------------
# Building
# --------------------------------------------------------------------------------------------------------------------
## Build binary
build:
	@echo "$(OK_COLOR)==> Building Binary$(NO_COLOR)\n"
	go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/${APP_NAME} cmd/${BINARY}/main.go

# --------------------------------------------------------------------------------------------------------------------
# Testing
# --------------------------------------------------------------------------------------------------------------------
## Run test
test:
	@printf "$(OK_COLOR)==> Running unit tests$(NO_COLOR)\n"
	@test -s $(GOPATH)/bin/overalls || go get -u github.com/go-playground/overalls
	@$(GOPATH)/bin/overalls -project=${IMPORT_PATH} -covermode=atomic -- -race

## -- Working with Docker --

# --------------------------------------------------------------------------------------------------------------------
# Docker
# --------------------------------------------------------------------------------------------------------------------
## Build docker image with name travels-budget
docker-init:
	@printf "$(OK_COLOR)==> Building docker image ${APP_NAME}$(NO_COLOR)\n"
	@docker image build -t ${APP_NAME} .

## Run tests using docker [travels-budget]
docker-test:
	@printf "$(OK_COLOR)==> Running test using image ${APP_NAME}$(NO_COLOR)\n"
	@docker run --rm -v $(PWD):/go/src/github.com/dohernandez/travels-budget ${APP_NAME} make deps && make test

## -- Running algorithm --
# --------------------------------------------------------------------------------------------------------------------
# Solution Space
# --------------------------------------------------------------------------------------------------------------------
## Runs the algorithm using the docker image [travels-budget]
itinerary-planner:
	@printf "$(OK_COLOR)==> Running itinerary-planner docker image ${APP_NAME}$(NO_COLOR)\n"
	@docker run --rm -v ${ACTIVITIES}:/resources/activities/activities.json ${APP_NAME} ${APP_NAME} personalized -b ${BUDGET} -d ${DAYS}

## Demo - Runs the algorithm using the values define in the problem using the docker image [travels-budget]; alias "travels-budget personalized -b 680 -d 2" using the file berlin.json
demo:
	@printf "$(OK_COLOR)==> Running demo docker image ${APP_NAME}$(NO_COLOR)\n"
	@docker run --rm ${APP_NAME} ${APP_NAME} personalized -b 680 -d 2
.PHONY: run clean deps build test docker-init docker-test demo itinerary-planner


# --------------------------------------------------------------------------------------------------------------------
# Help
# --------------------------------------------------------------------------------------------------------------------
.DEFAULT_GOAL := help
HELP_WIDTH="      "
help:
	@printf "\n";
	@printf "Base commands to work/dev the tool to inspire travelers to seek out the memorable things to do.\n";
	@awk '{ \
			if ($$0 ~ /^.PHONY: [a-zA-Z\-\_0-9]+$$/) { \
				helpCommand = substr($$0, index($$0, ":") + 2); \
				if (helpMessage) { \
					printf "  \033[32m%-20s\033[0m %s\n", \
						helpCommand, helpMessage; \
					helpMessage = ""; \
				} \
			} else if ($$0 ~ /^[a-zA-Z\-\_0-9.]+:/) { \
				helpCommand = substr($$0, 0, index($$0, ":")); \
				if (helpMessage) { \
					printf "  \033[32m%-20s\033[0m %s\n", \
						helpCommand, helpMessage; \
					helpMessage = ""; \
				} \
			} else if ($$0 ~ /^##/) { \
				if (helpMessage) { \
					helpMessage = helpMessage"\n"${HELP_WIDTH}substr($$0, 3); \
				} else { \
					helpMessage = substr($$0, 3); \
				} \
			} else { \
				if (helpMessage) { \
					print "\n"${HELP_WIDTH}helpMessage"\n" \
				} \
				helpMessage = ""; \
			} \
		}' \
		$(MAKEFILE_LIST)
	@printf "\n";
	@printf "Usage:\n";
	@printf ${HELP_WIDTH}"make itinerary-planner ACTIVITIES=<filepath> BUDGET=<budget> DAYS=<days>\n";
