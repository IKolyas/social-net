FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git

WORKDIR /app

COPY ./app/go.mod .
COPY ./app/go.sum .

RUN go mod download

COPY ./app .

# RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

ENTRYPOINT [ "tail", "-f" ,"/dev/null" ]

#FOR PRODUCTION
# FROM scratch

# WORKDIR /build

# COPY --from=builder /app/main /build/main

# ENTRYPOINT ["/build/main"]