.PHONY: generate
generate:
	@swagger generate client -f swagger.yaml -r SPDX

.PHONY: clean
clean:
	@rm -fr client
	@rm -fr models
