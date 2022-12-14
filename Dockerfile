FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/EricZapater/clashapi
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/clashapi cmd/api/main.go

FROM scratch
COPY --from=build /go/bin/clashapi /go/bin/clashapi
ENTRYPOINT ["/go/bin/clashapi"]