package common

import (
	"time"
)

type BasicEntity struct {
	ID       uint `gorm:"primary_key"`
	Ctime    time.Time
	Mtime    time.Time
	Dtime    time.Time
	Comment  string
	Priority int
	IsActive int
	IsDel    int
}
