build:
	protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=.  proto/user/user.proto
	# sudo docker build -t ecommerce-userservice .

run: 
	sudo docker rm -f userservice1
	sudo docker run -it -p 8805:8805 -e COUCHBASE_HOST=http://172.17.0.16 -e COUCHBASE_BUCKET=default -e COUCHBASE_USERNAME=Administrator -e COUCHBASE_PASSWORD=password --network="bridge" --name userservice1 ecommerce-userservice
	# sudo docker run -d -p 1111:8805  --network="bridge" --name userservice2 ecommerce-userservice

go:
	env COUCHBASE_HOST=http://localhost:8097/ COUCHBASE_BUCKET=default COUCHBASE_USERNAME=Administrator COUCHBASE_PASSWORD=password go run *.go

network:
	sudo docker network ls

clean:
	go clean -modcache