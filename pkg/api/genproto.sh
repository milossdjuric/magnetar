protoc --proto_path=./ \
        --go_out=. \
        --go_opt=paths=source_relative \
        --go-grpc_out=. \
        --go-grpc_opt=paths=source_relative \
        --go_opt=Mmagnetar.proto=github.com/c12s/magnetar/pkg/api \
        --go-grpc_opt=Mmagnetar.proto=github.com/c12s/magnetar/pkg/api \
        magnetar.proto