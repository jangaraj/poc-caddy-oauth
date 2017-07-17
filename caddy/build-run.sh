docker build -t caddy .
docker rm -f caddy | true
docker run -d \
  --name caddy \
  -e JWT_SECRET=secret \
  --net=host \
  caddy
