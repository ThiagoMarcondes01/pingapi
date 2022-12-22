FROM golang:1.13 as builder

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ping-api


FROM scratch
COPY --from=builder /app/server/ /server

ENTRYPOINT [ "/server" ]