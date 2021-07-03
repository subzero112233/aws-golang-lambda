package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"aws-golang-lambda/pkg/jwtparser"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	Unauthorized = errors.New("Unauthorized")
	secret       string
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := log.With().Logger()

	secret = os.Getenv("JWT_SECRET")
	if secret == "" {
		logger.Error().
			Msg("jwt secret environment varialbe is missing")
		os.Exit(1)
	}
}

func handler(event events.APIGatewayCustomAuthorizerRequestTypeRequest) (response events.APIGatewayCustomAuthorizerResponse, err error) {
	if event.Headers["Authorization"] == "" {
		event.Headers["Authorization"] = event.Headers["authorization"]
	}

	if !strings.HasPrefix(event.Headers["Authorization"], "Bearer") || len(event.Headers["Authorization"]) <= len("Bearer") {
		return response, Unauthorized
	}

	token := event.Headers["Authorization"][len("Bearer "):]

	claims, err := jwtparser.ValidateToken(token, secret)
	if err != nil {
		return response, Unauthorized
	}

	principal := fmt.Sprintf("user|%s", claims.Username)
	return CreatePolicy(principal), nil
}

func CreatePolicy(principal string) events.APIGatewayCustomAuthorizerResponse {
	// Create IAM policy that allows access
	return events.APIGatewayCustomAuthorizerResponse{
		PrincipalID: principal,
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   "Allow",
					Resource: []string{"*"},
				},
			},
		},
	}
}

func main() {
	lambda.Start(handler)
}
