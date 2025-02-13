package repo

import (
	"context"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/data"
)

type ProjectNodeRepo interface {
	FindAll(ctx context.Context) (list []*data.ProjectNode, err error)
}
