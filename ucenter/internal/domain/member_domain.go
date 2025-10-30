package domain

import (
	"common/msdb"
	"context"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberDomain struct {
	MemberRepo repo.MemberRepo
}

func NewMemberDomain(db *msdb.MsDB) *MemberDomain {
	return &MemberDomain{
		MemberRepo: dao.NewMemberDao(db),
	}
}

func (m *MemberDomain) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	// 涉及到数据库操作，此处省略
	mem, err := m.MemberRepo.FindByPhone(ctx, phone)
	if err != nil {
		logx.Error(ctx, "查询用户信息失败: %s", err.Error())
		return nil, err
	}
	return mem, nil
}
