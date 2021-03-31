all: clean snapshot

clean:
	@rm -rf dist

snapshot:
	goreleaser --snapshot --skip-publish --rm-dist

publish:
	goreleaser

run-docker: snapshot
	docker-compose -f build/package/docker-compose.dev.yml up

.PHONY: all build