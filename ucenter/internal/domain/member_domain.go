package domain

import (
	"common/msdb"
	"common/tools"
	"context"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberDomain struct {
	memberRepo repo.MemberRepo
}

func (m *MemberDomain) Register(context context.Context, phone string, password string, username string, country string, partner string, promotion string) error {
	mem := model.NewMember()

	// 对password进行md5加密
	_ = tools.Default(mem)
	salt, _, err := tools.Encode(password)
	if err != nil {
		return err
	}
	mem.Username = username
	mem.Country = country
	mem.Password = password
	mem.MobilePhone = phone
	mem.FillSuperPartner(partner)
	mem.PromotionCode = promotion
	mem.MemberLevel = model.GENERAL
	mem.Salt = salt
	mem.Avatar = "https://mszlu.oss-cn-shenzhen.aliyuncs.com/avatar/default.png"
	err = m.memberRepo.Save(context, mem)
	if err != nil {
		logx.Error(err)
		return err
	}
	return nil
}

func NewMemberDomain(db *msdb.MsDB) *MemberDomain {
	return &MemberDomain{
		memberRepo: dao.NewMemberDao(db),
	}
}

func (m *MemberDomain) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	// 涉及到数据库操作，此处省略
	mem, err := m.memberRepo.FindByPhone(ctx, phone)
	if err != nil {
		logx.Error(ctx, "查询用户信息失败: %s", err.Error())
		return nil, err
	}
	return mem, nil
}
