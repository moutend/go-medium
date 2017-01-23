.PHONY: test

fvl: fmt vet lint

fmt:
	@go fmt

vet:
	@go vet

lint:
	@golint -min_confidence=0.1

test:
	@go test | tee -a log
