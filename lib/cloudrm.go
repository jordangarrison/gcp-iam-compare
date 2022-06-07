package lib

import (
	"context"

	rm "google.golang.org/api/cloudresourcemanager/v1"
)

func NewRMService(ctx context.Context) (*rm.Service, error) {
	cloudresourcemanagerService, err := rm.NewService(ctx)
	return cloudresourcemanagerService, err
}
