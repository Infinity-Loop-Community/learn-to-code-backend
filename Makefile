.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build:
	go build ./...

.PHONY: test
test:
	go test ./...

.PHONY: sam-build
sam-build:
	sam build

.PHONY: deploy-dev
deploy-dev:
	sam build
	sam validate
	sam deploy --config-env dev

.PHONY: delete-dev
delete-dev:
	sam delete --region eu-central-1 --config-env dev

.PHONY: logs-dev-tail
logs-dev-tail:
	sam logs -n ParticipantPost --stack-name learn-to-code-backend --region eu-central-1 -t

.PHONY: clean-rebuild
clean-rebuild:
	rm -rf .aws-sam && make build && sam build

.PHONY: run-api-local
run-api-local:
	sam local start-api
