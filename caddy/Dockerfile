FROM golang:1.8 as build
RUN go get -d github.com/mholt/caddy/caddy \
  && go get -d github.com/tarent/loginsrv/caddy \
  && go get -d github.com/BTBurke/caddy-jwt \
  && cd /go/src/github.com/mholt/caddy/caddy \
  && sed -i '/This is where other plugins get plugged in (imported)/a     _ "github.com/tarent/loginsrv/caddy"' \
  /go/src/github.com/mholt/caddy/caddy/caddymain/run.go \
  && sed -i '/This is where other plugins get plugged in (imported)/a     _ "github.com/BTBurke/caddy-jwt"' \
  /go/src/github.com/mholt/caddy/caddy/caddymain/run.go \
  && CGO_ENABLED=0 ./build.bash
 
FROM ubuntu as cert
RUN apt-get update && apt-get install ca-certificates -y
 
FROM scratch
WORKDIR /
COPY --from=build /go/src/github.com/mholt/caddy/caddy/caddy /
COPY --from=cert /etc/ssl/certs /etc/ssl/certs
COPY login.tmpl /
COPY Caddyfile /
 
CMD ["/caddy", "-quic", "--conf", "/Caddyfile"]
