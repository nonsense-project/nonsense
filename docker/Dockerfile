# -- multistage docker build: stage #1: build stage
FROM golang:1.26.8-alpine AS build

RUN mkdir -p /go/src/github.com/nonsense-project/nonsense/

WORKDIR /go/src/github.com/nonsense-project/nonsense/

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache nonsensed dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o nonsensed .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/nonsense-project/nonsense/nonsensed /app/
COPY --from=build /go/src/github.com/nonsense-project/nonsense/infrastructure/config/sample-nonsensed.conf /app/

USER root
ENTRYPOINT [ "/sbin/tini", "--" ]
