package dynamodb

import (
	"context"
	"testing"

	aws_dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"	
)

func TestLocalClient(t *testing.T) {

	ctx := context.Background()

	client_uri := "aws://localhost?credentials=anon:"
	
	client, err := NewClient(ctx, client_uri)

	if err != nil {
		t.Fatalf("Failed to create new client, %v", err)
	}

	input := &aws_dynamodb.ListTablesInput{}

	_, err = client.ListTables(ctx, input)

	if err != nil {
		t.Fatalf("Failed to list tables, %v", err)
	}
	
}
	
