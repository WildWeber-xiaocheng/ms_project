package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"test.com/project-project/internal/data/pro"
	"test.com/project-project/internal/database"
	"test.com/project-project/internal/database/gorms"
)

type ProjectDao struct {
	conn *gorms.GormConn
}

func (p ProjectDao) FindProjectById(ctx context.Context, projectCode int64) (pj *pro.Project, err error) {
	err = p.conn.Session(ctx).Where("id = ?", projectCode).Find(&pj).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func (p ProjectDao) FindProjectMemberByPid(ctx context.Context, projectCode int64, page int64, pageSize int64) (list []*pro.ProjectMember, total int64, err error) {
	session := p.conn.Session(ctx)
	//和视频47 写的不一样，视频没有order limit offset
	err = session.Model(&pro.ProjectMember{}).
		Where("project_code=?", projectCode).
		Order("member_code asc").
		Limit(int(pageSize)).
		Offset(int((page - 1) * pageSize)).
		Find(&list).
		Error
	err = session.Model(&pro.ProjectMember{}).Where("project_code=?", projectCode).Count(&total).Error
	return
}

func (p ProjectDao) UpdateProject(ctx context.Context, proj *pro.Project) error {
	return p.conn.Session(ctx).Updates(&proj).Error
}

func (p ProjectDao) DeleteProjectCollect(ctx context.Context, memId int64, projectCode int64) error {
	return p.conn.Session(ctx).Where("member_code=? and project_code=?", memId, projectCode).Delete(&pro.ProjectCollection{}).Error
}

func (p ProjectDao) SaveProjectCollect(ctx context.Context, pc *pro.ProjectCollection) error {
	return p.conn.Session(ctx).Save(&pc).Error
}

func (p ProjectDao) UpdateDeletedProject(ctx context.Context, id int64, deleted bool) error {
	session := p.conn.Session(ctx)
	var err error
	if deleted {
		err = session.Model(&pro.Project{}).Where("id=?", id).Update("deleted", 1).Error
	} else {
		err = session.Model(&pro.Project{}).Where("id=?", id).Update("deleted", 0).Error
	}
	return err
}

func (p ProjectDao) FindProjectByPIdAndMemId(ctx context.Context, projectCode int64, memId int64) (*pro.ProjectAndMember, error) {
	var pms *pro.ProjectAndMember
	session := p.conn.Session(ctx)
	//sql和视频不一样
	sql := fmt.Sprintf("SELECT a.*, b.member_code, b.project_code, b.join_time, b.is_owner, b.authorize FROM ms_project a, ms_project_member b " +
		"WHERE a.id = b.project_code and b.member_code = ? and b.project_code = ? LIMIT 1")
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
