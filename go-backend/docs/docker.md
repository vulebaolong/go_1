<!-- build image -->
docker build -t img-go-backend:latest .

<!-- xoá image -->
docker image remove img-go-backend

<!-- xoá cache -->
docker builder prune -af

<!-- chạy image -->
docker run --name con-go-backend -d -p 3066:3069 --env-file .env img-go-backend:latest
docker run --name con-go-backend -d -p 3066:3069 --env-file .env.production img-go-backend:latest

docker logs -f con-go-backend

docker container remove con-go-backend
docker container stop con-go-backend
docker container start con-go-backend
docker container restart con-go-backend

docker container list
docker image list
docker network list

docker network create go-network

docker compose --env-file .env.production up -d
docker compose down