package repo

import (
	"context"
	"test.com/project-project/internal/data/pro"
	"test.com/project-project/internal/database"
)

type ProjectRepo interface {
	FindProjectByMemId(ctx context.Context, memId int64, condition string, page int64, size int64) ([]*pro.ProjectAndMember, int64, error)
	FindCollectProjectByMemId(ctx context.Context, memId int64, page int64, size int64) ([]*pro.ProjectAndMember, int64, error)
	SaveProject(ctx context.Context, conn database.DbConn, pr *pro.Project) error
	SaveProjectMember(ctx context.Context, conn database.DbConn, pm *pro.ProjectMember) error
	FindProjectByPIdAndMemId(ctx context.Context, projectCode int64, memId int64) (*pro.ProjectAndMember, error)
	FindCollectByPidAndMemId(ctx context.Context, projectCode int64, memId int64) (bool, error)
	UpdateDeletedProject(ctx context.Context, id int64, deleted bool) error
	SaveProjectCollect(ctx context.Context, pc *pro.ProjectCollection) error
	DeleteProjectCollect(ctx context.Context, memId int64, projectCode int64) error
	UpdateProject(ctx context.Context, proj *pro.Project) error
	FindProjectMemberByPid(ctx context.Context, projectCode int64, page int64, pageSize int64) (list []*pro.ProjectMember, total int64, err error)
	FindProjectById(ctx context.Context, projectCode int64) (pj *pro.Project, err error)
}

type ProjectTemplateRepo interface {
	FindProjectTemplateSystem(ctx context.Context, page int64, size int64) ([]pro.ProjectTemplate, int64, error)
	FindProjectTemplateCustom(ctx context.Context, memId int64, organizationCode int64, page int64, size int64) ([]pro.ProjectTemplate, int64, error)
	FindProjectTemplateAll(ctx context.Context, organizationCode int64, page int64, size int64) ([]pro.ProjectTemplate, int64, error)
}
