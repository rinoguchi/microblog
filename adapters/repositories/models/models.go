package repositories

import (
	"github.com/uptrace/bun"
	"time"
)

type DbComment struct {
	bun.BaseModel `bun:"table:comment,alias:c"`

	Id        *int64     `bun:"id,pk"`
	Text      string     `bun:"text"`
	CreatedAt *time.Time `bun:"created_at"`
	UpdatedAt *time.Time `bun:"updated_at"`
}
