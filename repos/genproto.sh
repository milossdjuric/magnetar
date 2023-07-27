protoc --proto_path=./ \
        --go_out=./ \
        --go_opt=paths=source_relative \
        --go_opt=Mnode.proto=github.com/c12s/magnetar/repos \
        node.proto