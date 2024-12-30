package dao

import (
	"context"
	"test.com/project-project/internal/data/pro"
	"test.com/project-project/internal/database/gorms"
)

type ProjectDao struct {
	conn *gorms.GormConn
}

func (p ProjectDao) FindProjectByMemId(ctx context.Context, memId int64, page int64, size int64) ([]*pro.ProjectAndMember, int64, error) {
	session := p.conn.Session(ctx)
	index := (page - 1) * size
	db := session.Raw("select * from ms_project a, ms_project_member b where a.id=b.project_code and b.member_code=? limit ?,?", memId, index, size)
	var mp []*pro.ProjectAndMember
	err := db.Scan(&mp).Error
	var total int64
	session.Model(&pro.ProjectMember{}).Where("member_code=?", memId).Count(&total)
	return mp, total, err
}

func NewProjectDao() *ProjectDao {
	return &ProjectDao{
		conn: gorms.New(),
	}
}
