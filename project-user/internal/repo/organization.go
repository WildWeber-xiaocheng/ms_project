package repo

import (
	"context"
	"github.com/WildWeber-xiaocheng/ms_project/project-user/internal/data/organization"
	"github.com/WildWeber-xiaocheng/ms_project/project-user/internal/database"
)

type OrganizationRepo interface {
	SaveOrganization(conn database.DbConn, ctx context.Context, org *organization.Organization) error
	FindOrganizationByMemId(ctx context.Context, memId int64) ([]*organization.Organization, error)
}
