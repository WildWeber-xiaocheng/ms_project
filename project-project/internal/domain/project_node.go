package domain

import (
	"context"
	"test.com/project-common/errs"
	"test.com/project-project/internal/dao"
	"test.com/project-project/internal/data"
	"test.com/project-project/internal/repo"
	"test.com/project-project/pkg/model"
)

type ProjectNodeDomain struct {
	projectNodeRepo repo.ProjectNodeRepo
}

func (d *ProjectNodeDomain) TreeList() ([]*data.ProjectNodeTree, *errs.BError) {
	//node表都查出来，转换成treelist结构
	nodes, err := d.projectNodeRepo.FindAll(context.Background())
	if err != nil {
		return nil, model.DBError
	}
	treeList := data.ToNodeTreeList(nodes)
	return treeList, nil
}

func (d *ProjectNodeDomain) AllNodeList() ([]*data.ProjectNode, *errs.BError) {
	nodes, err := d.projectNodeRepo.FindAll(context.Background())
	if err != nil {
		return nil, model.DBError
	}
	return nodes, nil
}

func NewProjectNodeDomain() *ProjectNodeDomain {
	return &ProjectNodeDomain{
		projectNodeRepo: dao.NewProjectNodeDao(),
	}
}
