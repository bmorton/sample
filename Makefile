.PHONY: help $(MAKECMDGOALS)
.DEFAULT_GOAL := snapshot

snapshot:
	goreleaser release --snapshot --rm-dist

clean:
	rm -rf dist/ build/
