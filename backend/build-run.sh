docker build -t backend .
docker rm -f backend
docker run --name backend -d -p 8081:8081 backend
