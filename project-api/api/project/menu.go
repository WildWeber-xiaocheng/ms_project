package project

import (
	"context"
	"github.com/WildWeber-xiaocheng/ms_project/project-api/pkg/model"
	common "github.com/WildWeber-xiaocheng/ms_project/project-common"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/errs"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/menu"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"time"
)

type HandlerMenu struct {
}

func (d *HandlerMenu) menuList(c *gin.Context) {
	result := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := MenuServiceClient.MenuList(ctx, &menu.MenuReqMessage{})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	var list []*model.Menu
	copier.Copy(&list, res.List)
	if list == nil {
		list = []*model.Menu{}
	}
	c.JSON(http.StatusOK, result.Success(list))
}

func NewMenu() *HandlerMenu {
	return &HandlerMenu{}
}
