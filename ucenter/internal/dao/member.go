package dao

import (
	"common/msdb"
	"common/msdb/gorms"
	"context"
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

func (m *MemberDao) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	session := m.conn.Session(ctx)
	mem := &model.Member{}
	err := session.Model(&model.Member{}).Where("mobile_phone = ?", phone).Take(mem).Error
	if err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return mem, nil
}
