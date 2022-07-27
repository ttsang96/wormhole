
## VeBridge Fork from Wormhole dev.v2 ## 

# Develop
**generate .proto to go file**

- brew install protobuf

- go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

- export PATH="$PATH:$(go env GOPATH)/bin"

- protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto


- git clone https://github.com/googleapis/googleapis.git       

- cp googleapis/google > google


'''

Run a node 
./node node --unsafeDevMode --guardianKey key --nodeName node1

'''