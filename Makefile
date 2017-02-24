install:
				@go get github.com/revel/revel
				@go get github.com/revel/cmd/revel
				@go get golang.org/x/oauth2

server:
				@revel run AuthKeyPush
