package store

import (
	"thinkdecideact/src/api"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type StoreController struct {
	DB      *gorm.DB
	Service StoreService
}

func NewStoreController(db *gorm.DB, service StoreService) StoreController {
	return StoreController{DB: db, Service: service}
}

func (c *StoreController) GetList(ctx iris.Context) {
	var pageCondition PageConditionDTO
	if err := ctx.ReadBody(&pageCondition); err != nil {
		api.Failure(ctx, err.Error())
		return
	}
	// TODO: In real situations, it is necessary to check the correctness of fields in pageCondition
	if pageCondition.PageIndex <= 0 {
		// Here, the first page starts from 0, not from 1.
		pageCondition.PageIndex = 0
	}
	if pageCondition.RowCountPerPage <= 0 {
		pageCondition.RowCountPerPage = 5
	}

	result, err := c.Service.GetManyByPage(pageCondition)
	if err != nil {
		api.Failure(ctx, err.Error())
		return
	}
	api.Success(ctx, api.SUCCESS_MSG, result)
}

func (c *StoreController) GetDetail(ctx iris.Context) {
	id := ctx.URLParamIntDefault("id", 0)
	if id <= 0 {
		api.Failure(ctx, "Invalid id")
		return
	}
	result, err := c.Service.GetOne(id)
	if err != nil {
		api.Failure(ctx, err.Error())
		return
	}
	api.Success(ctx, api.SUCCESS_MSG, result)
}

func (c *StoreController) Create(ctx iris.Context) {
	var storeEntity StoreEntity
	if err := ctx.ReadBody(&storeEntity); err != nil {
		api.Failure(ctx, err.Error())
		return
	}
	// TODO: In real situations, it is necessary to check the correctness of fields in storeEntity

	createdId, err := c.Service.Create(storeEntity)
	if err != nil {
		api.Failure(ctx, err.Error())
		return
	}
	api.Success(ctx, "Created successfully", map[string]interface{}{
		"createdId": createdId,
	})
}

func (c *StoreController) Delete(ctx iris.Context) {
	// id := ctx.PostValueIntDefault("id", 0)
	var deleteDTO DeleteDTO
	if err := ctx.ReadBody(&deleteDTO); err != nil {
		api.Failure(ctx, err.Error())
		return
	}
	// TODO: In real situations, it is necessary to check the correctness of fields in deleteDTO

	if err := c.Service.Delete(deleteDTO.ID); err != nil {
		api.Failure(ctx, err.Error())
		return
	}
	api.Success(ctx, "Deleted successfully")
}

func (c *StoreController) Update(ctx iris.Context) {
	// There are four solutions getting paramenter from the request.
	//
	// solution 1: the conventional way is to get one parameter at a time from the request: ctx.PostValue
	// name := ctx.PostValueDefault("name", "")
	// address := ctx.PostValueDefault("address", "")
	// if name == "" || address == "" {
	// 	ApiFailure(ctx, "name or address cannot be empty.")
	// 	return
	// }
	//
	// solution 2: store all the parameters into a map
	// var attrs map[string]interface{}
	// if err := ctx.ReadBody(&attrs); err != nil {
	// 	api.Failure(ctx, err.Error())
	// 	return
	// }
	//
	// solution 3: store all the parameters into an entity struct (Update all the fields of a record)
	// var storeEntity StoreEntity
	// if err := ctx.ReadBody(&storeEntity); err != nil {
	// 	api.Failure(ctx, err.Error())
	// 	return
	// }
	//
	// solution 4: store all the parameters into a dto struct (Partially update one or more fields of a record)
	var updateDTO UpdateDTO
	if err := ctx.ReadBody(&updateDTO); err != nil {
		api.Failure(ctx, err.Error())
		return
	}
	// TODO: In real situations, it is necessary to check the correctness of fields in updateDTO

	if err := c.Service.PartialUpdate(updateDTO.ID, updateDTO); err != nil {
		api.Failure(ctx, err.Error())
		return
	}
	api.Success(ctx, "Updated successfully")
}
