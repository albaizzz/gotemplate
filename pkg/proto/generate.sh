 #bin/bash
 protoc -I user/ user/user.proto --go_out=plugins=grpc:user


 mv user/*.pb.go ../pb/