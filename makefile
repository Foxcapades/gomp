
.PHONY: gen
gen: bin/gomp-gen
	@bin/gomp-gen v1/configs/config.yml
#	@gofmt -s -w .

bin/gomp-gen: v1/cmd/gomp-gen/main.go
	@go build -o bin/gomp-gen $?