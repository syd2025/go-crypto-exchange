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

func (m *MemberDao) UpdateLoginCount(ctx context.Context, id int64, step int) {
	session := m.conn.Session(ctx)
	err := session.Exec("update member set login_count = login_count + ? where id = ?", step, id).Error
	if err != nil {

	}
}

func (m *MemberDao) Save(ctx context.Context, mem *model.Member) error {
	session := m.conn.Session(ctx)
	err := session.Save(mem).Error
	return err
}
