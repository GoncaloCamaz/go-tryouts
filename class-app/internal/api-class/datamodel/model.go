package datamodel

import (
	"github.com/uptrace/bun"
	"time"
)

type Class struct {
	bun.BaseModel `bun:"table:class,alias:class"`

	ID      int64     `bun:"id,pk,autoincrement"`
	Number  int64     `bun:",notnull,type:int"`
	Year    string    `bun:",notnull,type:varchar(256)"`
	Created time.Time `bun:",notnull,default:current_timestamp"`
	Updated time.Time `bun:",notnull,default:current_timestamp"`
}
