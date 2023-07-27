protoc --proto_path=./ \
        --go_out=./ \
        --go_opt=paths=source_relative \
        --go_opt=Mregistration.proto=github.com/c12s/magnetar/handlers \
        registration.proto