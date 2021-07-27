FROM golang:1.16-alpine

RUN go get github.com/cespare/reflex

ENTRYPOINT ["reflex", "-c", "reflex.conf"]