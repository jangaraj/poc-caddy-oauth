docker build -t backend .
docker rm -f backend
docker run -d \
  --name backend \
  backend
