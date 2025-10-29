package tran

import "common/msdb"

type Transaction interface {
	Action(func(conn msdb.DbConn) error) error
}
