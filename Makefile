$(GOPATH)/bin/easyjson:
	@echo "Installing easyjson ..."
	@$(GO) get -u github.com/mailru/easyjson/...

.PHONY: easyjson
easyjson: $(GOPATH)/bin/easyjson
		@echo "Running easyjson generation"
		@$(GOPATH)/bin/easyjson -all \
			./app/domain/attribute.go \
			./app/domain/campaign.go \
			./app/domain/target.go \
            ./app/domain/target_collection.go \
            ./app/domain/search_result.go \
			./app/domain/attribute_collection.go
