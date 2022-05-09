package store

import "thinkdecideact/src/newtime"

// All forms of DTO
// queryDTO, conditionDTO (conditions of select/delete/update SQL)
// resultDTO (results of select SQL)
// createDTO (data used for insert SQL)
// updateDTO (data used for update SQL)
// paramDTO, argDTO, attrDTO

type PageConditionDTO struct {
	RowCountPerPage int `json:"rowCountPerPage" url:"rowCountPerPage" form:"rowCountPerPage"`
	PageIndex       int `json:"pageIndex" url:"pageIndex" form:"pageIndex"`
}

type ResultDTO struct {
	ID       uint           `json:"id"`
	Ctime    newtime.MyTime `json:"ctime"`
	Mtime    newtime.MyTime `json:"mtime"`
	Priority int            `json:"priority"`
	IsActive int            `json:"is_active"`
	Name     string         `json:"name"`
	Address  string         `json:"address"`
}

type CreateDTO struct {
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
}

type DeleteDTO struct {
	ID uint `json:"id" form:"id"`
}

type UpdateDTO struct {
	ID      uint   `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
}

func (_self StoreEntity) Serializer() ResultDTO {
	return ResultDTO{
		ID:       _self.ID,
		Ctime:    _self.Ctime,
		Mtime:    _self.Mtime,
		Priority: _self.Priority,
		IsActive: _self.IsActive,
		Name:     _self.Name,
		Address:  _self.Address,
	}
}
