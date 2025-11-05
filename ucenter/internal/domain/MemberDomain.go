package domain

import (
	"context"
	"errors"
	"mscoin-common/msdb"
	"mscoin-common/tools"
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

func (d *MemberDomain) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	// 涉及到数据库查询，需要使用事务
	mem, err := d.MemberRepo.FindByPhone(ctx, phone)
	if err != nil {
		logx.Error(err)
		return nil, errors.New("数据库异常")
	}
	return mem, nil
}

func (d *MemberDomain) Register(
	ctx context.Context,
	phone string,
	password string,
	username string,
	country string,
	partner string,
	promotion string,
) error {
	mem := model.NewMember()
	// 密码加密  md5 加盐
	_ = tools.Default(mem)
	salt, pwd := tools.Encode(password, nil)
	mem.Username = username
	mem.Country = country
	mem.Password = pwd
	mem.MobilePhone = phone
	mem.FillSuperPartner(partner)
	mem.PromotionCode = promotion
	mem.MemberLevel = model.GENERAL
	mem.Salt = salt
	mem.Avatar = "https://mscoin-1258344699.cos.ap-guangzhou.myqcloud.com/avatar/default.png"
	err := d.MemberRepo.Save(ctx, mem)
	if err != nil {
		logx.Error(err)
		return errors.New("数据库异常")
	}
	return nil
}
