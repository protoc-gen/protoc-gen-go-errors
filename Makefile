
.PHONY: init
# init env
init:
	go mod tidy
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.3

.PHONY: errors
# generate errors
errors:
	protoc --proto_path=./errors \
		   --proto_path=./third_party \
		   --go_out=paths=source_relative:./errors \
		   ./errors/*.proto

.PHONY: example
# generate example
example:
	go install . && \
	protoc --proto_path=. \
		   --go_out=paths=source_relative:. \
		   --go-errors_out=paths=source_relative:. \
		   ./example/*.proto

.PHONY: all
# generate all
all:
	make errors;
	go mod tidy;
