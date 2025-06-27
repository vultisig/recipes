default:
	 docker run --rm  -v "$(PWD)":/workspace -w /workspace --entrypoint /bin/sh bufbuild/buf "./build.sh"

gen-mocks:
	mockery
