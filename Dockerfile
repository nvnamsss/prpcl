FROM golang as builder
WORKDIR /app

COPY src/go.mod src/go.sum /app/

RUN go env -w GOPROXY=https://proxy.golang.org,direct && \
    go env -w GOSUMDB=off && go mod download

COPY . /app/

RUN make build
RUN mkdir build && cp -p ./src/cmd/prpcl ./build

FROM alpine:latest

RUN apk add tzdata ca-certificates
WORKDIR /app
COPY --from=builder /app/build /app

EXPOSE 8080

CMD ["./prpcl"]