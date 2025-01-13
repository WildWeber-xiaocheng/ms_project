package repo

import (
	"context"
	"test.com/project-project/internal/data/task"
	"test.com/project-project/internal/database"
)

type TaskStagesTemplateRepo interface {
	FindInProTemIds(ctx context.Context, ids []int) ([]task.MsTaskStagesTemplate, error)
	FindByProjectTemplatedId(ctx context.Context, projectTemplateCode int) (list []*task.MsTaskStagesTemplate, err error)
}

type TaskStagesRepo interface {
	SaveTaskStages(ctx context.Context, conn database.DbConn, stages *task.TaskStages) error
}
