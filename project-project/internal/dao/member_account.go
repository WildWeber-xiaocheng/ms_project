package dao

import (
	"context"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/data"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/database/gorms"
	"gorm.io/gorm"
)

type MemberAccountDao struct {
	conn *gorms.GormConn
}

func (m *MemberAccountDao) FindList(ctx context.Context, condition string, organizationCode int64, departmentCode int64, page int64, pageSize int64) (list []*data.MemberAccount, total int64, err error) {
	session := m.conn.Session(ctx)
	offset := (page - 1) * pageSize
	err = session.Model(&data.MemberAccount{}).
		Where("organization_code=?", organizationCode).
		Where(condition).Limit(int(pageSize)).Offset(int(offset)).Find(&list).Error
	err = session.Model(&data.MemberAccount{}).
		Where("organization_code=?", organizationCode).
		Where(condition).Count(&total).Error
	return
}

func (m *MemberAccountDao) FindByMemberId(ctx context.Context, memId int64) (ma *data.MemberAccount, err error) {
	session := m.conn.Session(ctx)
	err = session.Where("member_code=?", memId).Take(&ma).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func NewMemberAccountDao() *MemberAccountDao {
	return &MemberAccountDao{
		conn: gorms.New(),
	}
}
