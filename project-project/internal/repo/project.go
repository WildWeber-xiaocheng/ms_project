package repo

import (
	"context"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/data"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/database"
)

type ProjectRepo interface {
	FindProjectByMemId(ctx context.Context, memId int64, condition string, page int64, size int64) ([]*data.ProjectAndMember, int64, error)
	FindCollectProjectByMemId(ctx context.Context, memId int64, page int64, size int64) ([]*data.ProjectAndMember, int64, error)
	SaveProject(ctx context.Context, conn database.DbConn, pr *data.Project) error
	SaveProjectMember(ctx context.Context, conn database.DbConn, pm *data.ProjectMember) error
	FindProjectByPIdAndMemId(ctx context.Context, projectCode int64, memId int64) (*data.ProjectAndMember, error)
	FindCollectByPidAndMemId(ctx context.Context, projectCode int64, memId int64) (bool, error)
	UpdateDeletedProject(ctx context.Context, id int64, deleted bool) error
	SaveProjectCollect(ctx context.Context, pc *data.ProjectCollection) error
	DeleteProjectCollect(ctx context.Context, memId int64, projectCode int64) error
	UpdateProject(ctx context.Context, proj *data.Project) error
	FindProjectMemberByPid(ctx context.Context, projectCode int64, page int64, pageSize int64) (list []*data.ProjectMember, total int64, err error)
	FindProjectById(ctx context.Context, projectCode int64) (pj *data.Project, err error)
	FindProjectByIds(ctx context.Context, pids []int64) (list []*data.Project, err error)
}

type ProjectTemplateRepo interface {
	FindProjectTemplateSystem(ctx context.Context, page int64, size int64) ([]data.ProjectTemplate, int64, error)
	FindProjectTemplateCustom(ctx context.Context, memId int64, organizationCode int64, page int64, size int64) ([]data.ProjectTemplate, int64, error)
	FindProjectTemplateAll(ctx context.Context, organizationCode int64, page int64, size int64) ([]data.ProjectTemplate, int64, error)
}
