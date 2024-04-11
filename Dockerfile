FROM golang:1.22-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY ["./go.mod", "/go.sum", "./"]
RUN go mod download

#BUILD
COPY . .

RUN go build -o ./bin/currency-task cmd/currency-task/main.go
RUN go build -o ./bin/migrator cmd/migrator/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/currency-task /
COPY --from=builder /usr/local/src/config.json /
COPY --from=builder /usr/local/src/migrations /migrations
COPY --from=builder /usr/local/src/bin/migrator /


EXPOSE 8000

ENTRYPOINT ["sh", "-c", "sleep 5 && ./migrator && ./currency-task"]