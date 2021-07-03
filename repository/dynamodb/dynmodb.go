package ddb

import (
	"fmt"

	"aws-golang-lambda/entity"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

const (
	table = "users-table"
)

type DynamoDB struct {
	dynamodb dynamodbiface.DynamoDBAPI
}

func NewDynamoDB(dynamodbClient dynamodbiface.DynamoDBAPI) *DynamoDB {
	return &DynamoDB{
		dynamodb: dynamodbClient,
	}
}

func (d *DynamoDB) GetUser(username string) (details entity.User, err error) {
	result, err := d.dynamodb.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	})
	if err != nil {
		return details, fmt.Errorf("failed to get item from dynamodb: %s", err)
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &details)
	if err != nil {
		return details, fmt.Errorf("failed to unmarshal results with error: %w", err)
	}

	return details, nil
}

func (d *DynamoDB) AddUser(details entity.User) error {
	_, err := d.dynamodb.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(details.Username),
			},
			"password": {
				S: aws.String(details.Password),
			},
			"address": {
				S: aws.String(details.Address),
			},
			"first_name": {
				S: aws.String(details.FirstName),
			},
			"last_name": {
				S: aws.String(details.LastName),
			},
		},
		TableName: aws.String(table),
	})

	if err != nil {
		return fmt.Errorf("failed to insert to dynamodb with error: %v", err)
	}

	return nil
}
