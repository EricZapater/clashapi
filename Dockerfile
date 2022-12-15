FROM golang:alpine AS build

RUN apk add --update git
RUN apk --no-cache add tzdata
WORKDIR /go/src/github.com/EricZapater/clashapi
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/clashapi main.go

FROM scratch
COPY --from=build /go/bin/clashapi /go/bin/clashapi
ENTRYPOINT ["/go/bin/clashapi"]