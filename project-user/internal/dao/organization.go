package dao

import (
	"context"
	"test.com/project-user/internal/data/organization"
	"test.com/project-user/internal/database/gorms"
)

type OrganizationDao struct {
	conn *gorms.GormConn
}

func NewOrganizationDao() *OrganizationDao {
	return &OrganizationDao{
		conn: gorms.New(),
	}
}

func (o *OrganizationDao) SaveOrganization(ctx context.Context, org *organization.Organization) error {
	err := o.conn.Session(ctx).Create(org).Error
	return err
}
