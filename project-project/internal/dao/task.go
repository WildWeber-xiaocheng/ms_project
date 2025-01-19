package dao

import (
	"context"
	"gorm.io/gorm"
	"test.com/project-project/internal/data"
	"test.com/project-project/internal/database/gorms"
)

type TaskDao struct {
	conn *gorms.GormConn
}

func (t TaskDao) FindTaskMemberByTaskId(ctx context.Context, taskCode int64, memberId int64) (task *data.TaskMember, err error) {
	session := t.conn.Session(ctx)
	err = session.
		Where("task_code=? and member_code=?", taskCode, memberId).
		Limit(1).
		Find(&task).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func (t TaskDao) FindTaskByStageCode(ctx context.Context, stageCode int) (list []*data.Task, err error) {
	session := t.conn.Session(ctx)
	err = session.Model(&data.Task{}).
		Where("stage_code = ? and deleted = 0", stageCode).
		Order("sort asc").
		Find(&list).
		Error
	return
}

func NewTaskDao() *TaskDao {
	return &TaskDao{
		conn: gorms.New(),
	}
}
