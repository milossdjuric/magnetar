protoc --proto_path=./ \
        --go_out=./ \
        --go_opt=paths=source_relative \
        --go_opt=Mmodel.proto=github.com/c12s/magnetar/pkg/proto \
        model.proto

protoc --proto_path=./ \
        --go_out=. \
        --go_opt=paths=source_relative \
        --go-grpc_out=. \
        --go-grpc_opt=paths=source_relative \
        --go_opt=Mmagnetar.proto=github.com/c12s/magnetar/pkg/proto \
        --go-grpc_opt=Mmagnetar.proto=github.com/c12s/magnetar/pkg/proto \
        -I ./model.proto \
        magnetar.proto