
.PHONY: gen
gen: bin/gomp-gen
	@bin/gomp-gen --skip-tests v1/configs/config.yml

bin/gomp-gen: v1/cmd/gomp-gen/main.go
	@go build -o bin/gomp-gen $?