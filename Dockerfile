FROM golang:1.11

ARG VERSION=dev
ARG USER=dohernandez

WORKDIR /go/src/github.com/dohernandez/travels-budget

COPY . .

RUN make run

RUN cp /go/src/github.com/dohernandez/travels-budget/bin/travels-budget /bin/travels-budget

COPY resources/activities /resources/activities