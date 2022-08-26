FROM golang:alpine as build
RUN apk add git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o sandwiches-after-dark

FROM scratch
COPY --from=build /app/sandwiches-after-dark /sandwiches-after-dark
ENTRYPOINT ["/sandwiches-after-dark"]
