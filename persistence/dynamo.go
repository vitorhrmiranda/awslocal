package persistence

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/vitorhrmiranda/go-awslocal/entity"
	"github.com/vitorhrmiranda/go-awslocal/values"
)

type DynamoDB struct {
	c *dynamodb.DynamoDB
}

func NewDynamoDB() (DynamoDB, error) {
	var sess *session.Session
	var err error

	impl := DynamoDB{}

	sess, err = session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("foo", "var", ""),
		Region:      aws.String(values.DEFAULT_REGION),
		Endpoint:    aws.String(values.DYNAMODB_AWS_ENDPOINT),
	})

	if err != nil {
		return impl, err
	}

	impl.c = dynamodb.New(sess)

	return impl, nil
}

func (db *DynamoDB) Create(usr *entity.User) error {

	item, err := dynamodbattribute.MarshalMap(usr)

	if err != nil {
		return fmt.Errorf("MarshalMap: %s", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(values.TABLE_NAME),
		Item:      item,
	}

	if _, err := db.c.PutItem(input); err != nil {
		return fmt.Errorf("PutItem: %s", err)
	}

	return nil
}
