package lib

import (
	"context"

	cloudResourceManager "google.golang.org/api/cloudresourcemanager/v1"
)

func NewRMService(ctx context.Context) (*cloudResourceManager.Service, error) {
	cloudresourcemanagerService, err := cloudResourceManager.NewService(ctx)
	return cloudresourcemanagerService, err
}
