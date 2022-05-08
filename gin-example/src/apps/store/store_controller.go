package store

import (
	"log"
	"strconv"
	"thinkdecideact/src/api"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StoreController struct {
	DB      *gorm.DB
	Service StoreService
}

func NewStoreController(db *gorm.DB, service StoreService) StoreController {
	return StoreController{DB: db, Service: service}
}

func (c *StoreController) GetList(ctx *gin.Context) {
	var pageCondition PageConditionDTO
	if err := ctx.ShouldBind(&pageCondition); err != nil {
		api.Failure(ctx, err.Error())
		return
	}
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

func (c *StoreController) GetDetail(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		api.Failure(ctx, err.Error())
		return
	}

	result, err := c.Service.GetOne(id)
	if err != nil {
		api.Failure(ctx, err.Error())
		return
	}
	api.Success(ctx, api.SUCCESS_MSG, result)
}

func (c *StoreController) Create(ctx *gin.Context) {
	var storeEntity StoreEntity
	if err := ctx.ShouldBind(&storeEntity); err != nil {
		api.Failure(ctx, err.Error())
		return
	}
	log.Println(storeEntity)
	// TODO: In real situations, it is necessary to check the correctness of fields in createDTO

	createdId, err := c.Service.Create(storeEntity)
	if err != nil {
		api.Failure(ctx, err.Error())
		return
	}
	api.Success(ctx, "Created successfully", map[string]interface{}{
		"createdId": createdId,
	})
}

func (c *StoreController) Delete(ctx *gin.Context) {
	// id, err := strconv.Atoi(ctx.PostForm("id"))
	var deleteDTO DeleteDTO
	if err := ctx.ShouldBind(&deleteDTO); err != nil {
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

func (c *StoreController) Update(ctx *gin.Context) {
	var updateDTO UpdateDTO
	if err := ctx.ShouldBind(&updateDTO); err != nil {
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
