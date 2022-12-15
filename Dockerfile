FROM golang:alpine AS build

RUN apk add --update git
RUN apk add --no-cache tzdata 

WORKDIR /go/src/github.com/EricZapater/clashapi
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/clashapi main.go

FROM scratch
COPY --from=build /go/bin/clashapi /go/bin/clashapi
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENV TZ=Europe/Madrid

ENTRYPOINT ["/go/bin/clashapi"]