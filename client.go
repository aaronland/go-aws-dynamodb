package dynamodb

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	
	"github.com/aaronland/go-aws-auth"
	"github.com/aws/aws-sdk-go-v2/aws"	
	aws_dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewClientWithURI(ctx context.Context, uri string) (*aws_dynamodb.Client, error) {
	return NewClient(ctx, uri)
}

func NewClient(ctx context.Context, uri string) (*aws_dynamodb.Client, error) {
	
	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}
	
	cfg, err := auth.NewConfig(ctx, uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create config, %w", err)
	}

	is_local := false
	
	q := u.Query()

	if q.Has("local") {

		v, err := strconv.ParseBool(q.Get("local"))

		if err != nil {
			return nil, fmt.Errorf("Invalid ?local= parameter, %w", err)
		}

		is_local = v
	}

	if is_local {

		cfg.Region = "localhost"
		
		cfg.EndpointResolver = aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000", SigningRegion: "localhost"}, nil
			})
		
	}
	
	client := aws_dynamodb.NewFromConfig(cfg)
	return client, nil
}
