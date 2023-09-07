REPO=blacktop
NAME=app-icon
NEXT_VERSION=$(shell svu patch)

.PHONY: release
release: ## Create a new release from the NEXT_VERSION
	@echo " > Creating Release ${NEXT_VERSION}"
	git tag -a ${NEXT_VERSION} -m "${NEXT_VERSION}"
	git push origin ${NEXT_VERSION}
	goreleaser --clean --timeout 60m

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help