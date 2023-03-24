run-proto:
	for f in proto/*/*.proto; do \
		protoc --go_out=. $$f; \
		protoc --go-grpc_out=. $$f; \
		echo compiled: $$f; \
	done