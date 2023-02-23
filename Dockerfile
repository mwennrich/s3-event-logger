FROM golang:1.20 AS builder
ENV GO111MODULE=on
ENV CGO_ENABLED=0

COPY / /work
WORKDIR /work

RUN make s3-event-logger

FROM scratch
COPY --from=builder /work/bin/s3-event-logger /s3-event-logger

USER 999
ENTRYPOINT ["/s3-event-logger"]
