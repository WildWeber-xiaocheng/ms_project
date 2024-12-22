package repo

import (
	"context"
	"test.com/project-user/internal/data/organization"
)

type OrganizationRepo interface {
	SaveOrganization(ctx context.Context, org *organization.Organization) error
}
