package repo

import (
	"context"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/data"
)

type MenuRepo interface {
	FindMenus(ctx context.Context) ([]*data.ProjectMenu, error)
}
