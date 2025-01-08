package dao

import (
	"context"
	"fmt"
	"test.com/project-project/internal/data/pro"
	"test.com/project-project/internal/database"
	"test.com/project-project/internal/database/gorms"
)

type ProjectDao struct {
	conn *gorms.GormConn
}

func (p ProjectDao) FindProjectByPIdAndMemId(ctx context.Context, projectCode int64, memId int64) (*pro.ProjectAndMember, error) {
	var pms *pro.ProjectAndMember
	session := p.conn.Session(ctx)
	//sql和视频不一样
	sql := fmt.Sprintf("SELECT * FROM ms_project a, ms_project_member b " +
		"WHERE a.id = b.project_code and b.member_code = ? and b.id = ? LIMIT 1")
	raw := session.Raw(sql, memId, projectCode)
	err := raw.Scan(&pms).Error
	return pms, err
}

func (p ProjectDao) FindCollectByPidAndMemId(ctx context.Context, projectCode int64, memId int64) (bool, error) {
	var count int64
	session := p.conn.Session(ctx)
	sql := fmt.Sprintf("select  count(*) from ms_project_collection where member_code=? and project_code = ?")
	raw := session.Raw(sql, memId, projectCode)
	err := raw.Scan(&count).Error
	return count > 0, err
}

func (p ProjectDao) SaveProject(ctx context.Context, conn database.DbConn, pr *pro.Project) error {
	p.conn = conn.(*gorms.GormConn)
	return p.conn.Tx(ctx).Save(&pr).Error
}

func (p ProjectDao) SaveProjectMember(ctx context.Context, conn database.DbConn, pm *pro.ProjectMember) error {
	p.conn = conn.(*gorms.GormConn)
	return p.conn.Tx(ctx).Save(&pm).Error
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
