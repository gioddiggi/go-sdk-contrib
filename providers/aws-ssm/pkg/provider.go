package awsssm

import "context"

const providerName = "AWS SSM"

type Provider struct{}

func NewProvider() (*Provider, error) {
	return NewProviderWithContext(context.Background())
}

func NewProviderWithContext(ctx context.Context) (*Provider, error) {
	return &Provider{}, nil
}
