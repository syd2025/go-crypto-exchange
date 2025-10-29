package gorms

import (
	"context"

	"gorm.io/gorm"
)

type GormConn struct {
	db *gorm.DB
	tx *gorm.DB
}

func (c *GormConn) Begin() {
	c.tx = c.db.Begin()
}

func New(db *gorm.DB) *GormConn {
	return &GormConn{
		db: db,
		tx: db,
	}
}

func (c *GormConn) Session(ctx context.Context) *gorm.DB {
	return c.db.Session(&gorm.Session{Context: ctx})
}

func (c *GormConn) Rollback() {
	if c.tx != nil {
		c.tx.Rollback()
	}
}

func (c *GormConn) Commit() {
	if c.tx != nil {
		c.tx.Commit()
	}
}

func (c *GormConn) Tx(ctx context.Context) *gorm.DB {
	return c.tx.WithContext(ctx)
}
