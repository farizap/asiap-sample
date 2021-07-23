FROM golang:1.16-alpine

RUN go get github.com/cespare/reflex

# RUN go mod download

ENTRYPOINT ["reflex", "-c", "reflex.conf"]