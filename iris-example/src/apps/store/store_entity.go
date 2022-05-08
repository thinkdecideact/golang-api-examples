package store

import "thinkdecideact/src/newtime"

type StoreEntity struct {
	ID       uint           `gorm:"primary_key" json:"id"`
	Ctime    newtime.MyTime `json:"ctime"`
	Mtime    newtime.MyTime `json:"mtime"`
	Comment  string         `json:"comment"`
	Priority int            `json:"priority"`
	IsActive int            `json:"is_active"`
	IsDel    int            `json:"is_del"`
	Name     string         `json:"name" form:"name"`
	Address  string         `json:"address" form:"address"`
}

// make a mapping of struct StoreEntity to table `tdar_store`
func (_self StoreEntity) TableName() string {
	return "tdar_store"
}

func (_self *StoreEntity) Enable() {
	_self.IsActive = 1
}

func (_self *StoreEntity) Disable() {
	_self.IsActive = 0
}
