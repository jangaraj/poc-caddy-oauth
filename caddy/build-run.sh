docker build -t caddy .
docker rm -f caddy | true
docker run --name caddy --link backend -e JWT_SECRET=secret -d -p 2015:2015 caddy
