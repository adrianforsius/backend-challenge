FROM golang:1.16 as builder

WORKDIR /builder

ADD ./ .


RUN go mod vendor
# We set CGO_ENABLED=0 because:
# https://stackoverflow.com/questions/36279253/go-compiled-binary-wont-run-in-an-alpine-docker-container-on-ubuntu-host
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o app .

FROM alpine:3.8

RUN apk update && apk add ca-certificates
WORKDIR /bin
EXPOSE 8082

COPY --from=builder /builder/app ./app

ENTRYPOINT ["./app"]
