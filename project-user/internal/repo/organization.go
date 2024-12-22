package repo

import "context"

type OrganizationRepo interface {
	SaveOrganization(ctx context.Context, org any) error
}
