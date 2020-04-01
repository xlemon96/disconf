package bean

import "time"

type BaseModel struct {
	ID       int64
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
}
