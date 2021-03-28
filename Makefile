generate:
	buf generate

lint:
	buf lint
	buf breaking --against 'https://github.com/abdybaevae/url-shortener.git#branch=master'

BUF_VERSION:=0.40.0

install:
	go install \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	curl -sSL \
    	"https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(shell uname -s)-$(shell uname -m)" \
    	-o "$(shell go env GOPATH)/bin/buf" && \
  	chmod +x "$(shell go env GOPATH)/bin/buf"

hot-tests:
	reflex -c reflex.conf

MOCKS_DESTINATION=mocks
.PHONY: mocks
mocks: pkg/services/algo/type.go pkg/repos/algo/type.go pkg/services/key/type.go pkg/repos/key/type.go pkg/services/link/type.go pkg/repos/link/type.go pkg/services/number/type.go pkg/repos/number/type.go
	@echo "Generating mocks..."
	@rm -rf $(MOCKS_DESTINATION)
	@for file in $^; do mockgen -source=$$file -destination=$(MOCKS_DESTINATION)/$$file; done

# .PHONY: migrate 
# migrate: 

# migrate create -ext sql -dir db/migrations -seq create_links_table
# migrate -database "postgres://cifer@localhost:5432/url_shortener?sslmode=disable" -path db/migrations up 3



# drop table algorithms;
# drop table numbers;
# drop table schema_migrations;