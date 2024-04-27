FROM golang:alpine3.19 as build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go test -v
RUN go build -o /bin/app

FROM alpine:3.16.2
COPY --from=build /bin/app /bin
CMD ["/bin/app"]