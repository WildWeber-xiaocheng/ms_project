package dao

import (
	"context"
	"fmt"
	"test.com/project-project/internal/data/pro"
	"test.com/project-project/internal/database/gorms"
)

type ProjectDao struct {
	conn *gorms.GormConn
}

func (p ProjectDao) FindCollectProjectByMemId(ctx context.Context, memId int64, page int64, size int64) ([]*pro.ProjectAndMember, int64, error) {
	session := p.conn.Session(ctx)
	index := (page - 1) * size
	sql := fmt.Sprintf("select * from ms_project where id in (select project_code from ms_project_collection where member_code=? ) order by sort limit ?,?")
	db := session.Raw(sql, memId, index, size)
	var mp []*pro.ProjectAndMember
	err := db.Scan(&mp).Error
	var total int64
	query := fmt.Sprintf("member_code=?")
	session.Model(&pro.ProjectCollection{}).Where(query, memId).Count(&total)
	return mp, total, err
}

func (p ProjectDao) FindProjectByMemId(ctx context.Context, memId int64, condition string, page int64, size int64) ([]*pro.ProjectAndMember, int64, error) {
	session := p.conn.Session(ctx)
	index := (page - 1) * size
	sql := fmt.Sprintf("select * from ms_project a, ms_project_member b where a.id=b.project_code and b.member_code=? %s order by sort limit ?,?", condition)
	db := session.Raw(sql, memId, index, size)
	var mp []*pro.ProjectAndMember
	err := db.Scan(&mp).Error
	var total int64
	query := fmt.Sprintf("select count(*) from ms_project a, ms_project_member b where a.id = b.project_code and b.member_code=? %s", condition)
	tx := session.Raw(query, memId)
	err = tx.Scan(&total).Error
	return mp, total, err
}

func NewProjectDao() *ProjectDao {
	return &ProjectDao{
		conn: gorms.New(),
	}
}
