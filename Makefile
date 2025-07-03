include .env
export

run:
	go run main.go

load-images:

docker load -i exchange1_amd64.tar
docker load -i exchange2_amd64.tar
docker load -i exchange3_amd64.tar

run-images:

docker run -p 40101:40101 --name exchange1-arm64 -d exchange3-arm64
docker run -p 40102:40102 --name exchange2-arm64 -d exchange2-arm64
docker run -p 40103:40103 --name exchange3-arm64 -d exchange3-arm64
