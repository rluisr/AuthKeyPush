install:
        @go get github.com/Masterminds/glide
        @go install github.com/Masterminds/glide
        @glide install

server:
        @revel run github.com/rluisr/AuthKeyPush