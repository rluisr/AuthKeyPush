install:
        @go install github.com/revel/revel
        @go install github.com/revel/cmd/revel

server:
        @revel run AuthKeyPush