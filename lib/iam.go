package lib

import (
	"context"

	iam "google.golang.org/api/iam/v1"
)

func NewService(ctx context.Context) (*iam.Service, error) {
	iamService, err := iam.NewService(ctx)
	return iamService, err
}
