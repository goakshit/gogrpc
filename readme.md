#### Overview

Sample GRPC flow replicating unix echo command. 

#### Command to compile proto file: 
```
protoc ./hello.proto --go_out=plugins=grpc:.
```

##### Run main.go(GRPC server) in one terminal and client.go(client) in another. 