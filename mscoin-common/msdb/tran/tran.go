package tran

import "mscoin-common/msdb"

type Transaction interface {
	Action(func(conn msdb.DbConn) error) error
}
