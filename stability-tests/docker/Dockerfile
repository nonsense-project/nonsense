ARG NONSENSED_IMAGE
ARG NONSENSEMINER_IMAGE

FROM ${NONSENSED_IMAGE} as nonsensed
FROM ${NONSENSEMINER_IMAGE} as nonsenseminer

FROM golang:1.26.8-alpine

RUN mkdir -p /go/src/github.com/nonsense-project/nonsense/

WORKDIR /go/src/github.com/nonsense-project/nonsense/

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=nonsensed /app/ /app/
COPY --from=nonsenseminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/nonsense-project/nonsense/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]
