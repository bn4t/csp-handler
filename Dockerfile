FROM golang:alpine as builder

RUN mkdir /build /out && apk add --no-cache git ca-certificates

WORKDIR /build
COPY . .
RUN go get && go build -o /out/csp-handler





FROM alpine:latest

RUN mkdir /app && addgroup -S csp-handler && adduser -S csp-handler -G csp-handler && apk add --no-cache ca-certificates
COPY --from=builder /out/csp-handler /app/csp-handler
USER csp-handler

ENTRYPOINT /app/csp-handler

