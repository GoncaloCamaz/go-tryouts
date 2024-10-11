package datamodel

import (
	ClassDataModel "class-app/internal/api-class/datamodel"
	"github.com/uptrace/bun"
	"time"
)

type Student struct {
	bun.BaseModel `bun:"table:student,alias:student"`

	ID      int       `bun:"id,pk,autoincrement"`
	ClassId int       `bun:",notnull,type:int"`
	Name    string    `bun:",notnull,type:varchar(64)"`
	Email   string    `bun:",notnull,type:varchar(64)"`
	Created time.Time `bun:",notnull,default:current_timestamp"`
	Updated time.Time `bun:",notnull,default:current_timestamp"`

	Class *ClassDataModel.Class `bun:"rel:belongs-to,join:class_id=id"`
}
