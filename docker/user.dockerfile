FROM golang:1.18.6-alpine3.16 as builder

ARG VERSION

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o app -ldflags "-X 'main.version=${VERSION}'" ./cmd/user-api

FROM alpine:3.17.0 as runner

WORKDIR /app
COPY --from=builder /src/app ./

CMD ["./app"]