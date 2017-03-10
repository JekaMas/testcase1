EASYJSON_STRUCTS := $(shell find ./app/domain/*.go  -not -name '*_easyjson.go' -not -name '*_test.go'  -not -name 'helpers.go' -not -name 'verifier.go')

$(GOPATH)/bin/easyjson:
	@echo "Installing easyjson ..."
	@$(GO) get -u github.com/mailru/easyjson/...

.PHONY: easyjson
easyjson: $(GOPATH)/bin/easyjson
		@echo "Running easyjson generation"
		@$(GOPATH)/bin/easyjson -all $(EASYJSON_STRUCTS)