package store

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(routerGroup *gin.RouterGroup, db *gorm.DB) {
	storeService := NewStoreService(db)
	storeController := StoreController{DB: db, Service: storeService}

	pathPrefix := "/store"

	// http://127.0.0.1:8080/api/store/getList
	routerGroup.GET(pathPrefix+"/getList", storeController.GetList)

	// http://127.0.0.1:8080/api/store/getDetail
	routerGroup.GET(pathPrefix+"/getDetail", storeController.GetDetail)

	// http://127.0.0.1:8080/api/store/create
	routerGroup.POST(pathPrefix+"/create", storeController.Create)

	// http://127.0.0.1:8080/api/store/delete
	routerGroup.POST(pathPrefix+"/delete", storeController.Delete)

	// http://127.0.0.1:8080/api/store/update
	routerGroup.POST(pathPrefix+"/update", storeController.Update)
}
