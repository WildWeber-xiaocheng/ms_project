package repo

import (
	"context"
	"test.com/project-user/internal/data/member"
)

type MemberRepo interface {
	GetMemberByEmail(ctx context.Context, email string) (bool, error)
	GetMemberByAccount(ctx context.Context, account string) (bool, error)
	GetMemberByMobile(ctx context.Context, mobile string) (bool, error)
	SaveMember(ctx context.Context, mem *member.Member) error
}
