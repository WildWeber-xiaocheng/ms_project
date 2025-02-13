package menu_service_v1

import (
	"context"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/errs"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/menu"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/dao"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/database/tran"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/domain"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/repo"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type MenuService struct {
	menu.UnimplementedMenuServiceServer
	cache       repo.Cache
	transaction tran.Transaction
	menuDomain  *domain.MenuDomain
}

func New() *MenuService {
	return &MenuService{
		cache:       dao.Rc,
		transaction: dao.NewTransaction(),
		menuDomain:  domain.NewMenuDomain(),
	}
}

func (m *MenuService) MenuList(context.Context, *menu.MenuReqMessage) (*menu.MenuResponseMessage, error) {
	treeList, err := m.menuDomain.MenuTreeList()
	if err != nil {
		zap.L().Error("MenuList error", zap.Error(err))
		return nil, errs.GrpcError(err)
	}
	var list []*menu.MenuMessage
	copier.Copy(&list, treeList)
	return &menu.MenuResponseMessage{List: list}, nil
}
