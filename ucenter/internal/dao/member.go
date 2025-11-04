package dao

import (
	"context"
	"mscoin-common/msdb"
	"mscoin-common/msdb/gorms"
	"ucenter/internal/model"

	"gorm.io/gorm"
)

type MemberDao struct {
	conn *gorms.GormConn
}

func NewMemberDao(db *msdb.MsDB) *MemberDao {
	return &MemberDao{
		conn: gorms.New(db.Conn),
	}
}

func (m MemberDao) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	session := m.conn.Session(ctx)
	mem := &model.Member{}
	err := session.Model(&model.Member{}).Where("mobile_phone = ?", phone).Limit(1).Take(&mem).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	return mem, err
}
