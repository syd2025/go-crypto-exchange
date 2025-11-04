package tran

import (
	"mscoin-common/msdb"
	"mscoin-common/msdb/gorms"

	"gorm.io/gorm"
)

type TransactionImpl struct {
	conn msdb.DbConn
}

func NewTransaction(db *gorm.DB) *TransactionImpl {
	return &TransactionImpl{
		conn: gorms.New(db),
	}
}

// 事务处理机制
func (t *TransactionImpl) Action(f func(conn msdb.DbConn) error) error {
	t.conn.Begin()
	err := f(t.conn)
	if err != nil {
		t.conn.Rollback()
		return err
	}
	t.conn.Commit()
	return nil
}
