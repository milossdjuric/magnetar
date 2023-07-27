protoc --proto_path=./ \
        --go_out=./ \
        --go_opt=paths=source_relative \
        --go_opt=Mmodel.proto=github.com/c12s/magnetar/pkg/marshallers/proto \
        model.proto