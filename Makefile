.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build:
	go build ./...

.PHONY: test
build:
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
