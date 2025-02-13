package repo

import (
	"context"
	"github.com/WildWeber-xiaocheng/ms_project/project-user/internal/data/member"
	"github.com/WildWeber-xiaocheng/ms_project/project-user/internal/database"
)

type MemberRepo interface {
	GetMemberByEmail(ctx context.Context, email string) (bool, error)
	GetMemberByAccount(ctx context.Context, account string) (bool, error)
	GetMemberByMobile(ctx context.Context, mobile string) (bool, error)
	SaveMember(conn database.DbConn, ctx context.Context, mem *member.Member) error
	FindMember(ctx context.Context, account string, pwd string) (mem *member.Member, err error)
	FindMemberById(ctx context.Context, id int64) (mem *member.Member, err error)
	FindMemberByIds(ctx context.Context, ids []int64) (list []*member.Member, err error)
}
