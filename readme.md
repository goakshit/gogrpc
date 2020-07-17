#### Overview

Simple GRPC Service to get complete fetch blogs from mongodb. 

#### Command to compile proto file: 
```
protoc ./blog.proto --go_out=plugins=grpc:.
```

##### Run main.go(GRPC server) in one terminal and client.go(client) in another. 