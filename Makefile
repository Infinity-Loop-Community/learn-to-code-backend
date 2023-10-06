.PHONY: build deploy-dev delete-dev logs-dev-tail

build:
	sam build

deploy-dev:
	sam build
	sam validate
	sam deploy --config-env dev

delete-dev:
	sam delete --region eu-central-1 --config-env dev

logs-dev-tail:
	sam logs -n ParticipantPost --stack-name learn-to-code-backend --region eu-central-1 -t