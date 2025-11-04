package domain

import (
	"context"
	"ucenter/internal/model"
)

type MemberDomain struct {
}

func (d *MemberDomain) FindByPhone(ctx context.Context, phone string) *model.Member {
	// 涉及到数据库查询，需要使用事务
	return nil
}
