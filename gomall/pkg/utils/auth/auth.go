package auth

import "context"

type AuthToken struct {
	Token string
}

func (c AuthToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": c.Token,
	}, nil
}

func (c AuthToken) RequireTransportSecurity() bool {
	return false
}
