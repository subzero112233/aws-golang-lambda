.DEFAULT_GOAL := build
BUCKET_NAME=resheftest

build:
	@go mod tidy
	@cd lambdas/users && GOOS=linux CGO_ENABLED=0 go build -o lambda_handler .
	@cd lambdas/authorizer && go build -o lambda_handler .

deploy:
	cd lambdas/users && ./deploy.sh ${BUCKET_NAME}
	cd lambdas/authorizer && ./deploy.sh ${BUCKET_NAME}

deploy-apig:
	cd cloudformation/apigateway && ./create_apigateway.sh

test:
	go test -v -coverprofile cover.out . && go tool cover -html=cover.out -o cover.html
