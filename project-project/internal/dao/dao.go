package dao

import (
	"errors"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/errs"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/database"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/database/gorms"
)

type TransactionImpl struct {
	conn database.DbConn
}

func (t *TransactionImpl) Action(f func(conn database.DbConn) error) error {
	t.conn.Begin()
	err := f(t.conn)
	var bErr *errs.BError
	if errors.Is(err, bErr) {
		bErr = err.(*errs.BError)
		if bErr != nil {
			t.conn.Rollback()
			return bErr
		} else {
			t.conn.Commit()
			return nil
		}
	}
	if err != nil {
		t.conn.Rollback()
		return err
	}
	t.conn.Commit()
	return nil
}

func NewTransaction() *TransactionImpl {
	return &TransactionImpl{
		conn: gorms.NewTran(),
	}
}
