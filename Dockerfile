FROM golang:alpine AS build

RUN apk add --update git
RUN apk add --no-cache tzdata
RUN ln -fs /usr/share/zoneinfo/Europe/Madrid /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata
WORKDIR /go/src/github.com/EricZapater/clashapi
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/clashapi main.go

FROM scratch
COPY --from=build /go/bin/clashapi /go/bin/clashapi
ENV TZ="Europe/Madrid"
ENTRYPOINT ["/go/bin/clashapi"]