package store

import (
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

func Router(r iris.Party, db *gorm.DB) {
	storeAPI := r.Party("/store")
	{
		storeAPI.Use(iris.Compression)

		storeService := NewStoreService(db)
		storeController := StoreController{DB: db, Service: storeService}

		// http://127.0.0.1:8080/api/store/getList
		storeAPI.Get("/getList", storeController.GetList)

		// http://127.0.0.1:8080/api/store/getDetail
		storeAPI.Get("/getDetail", storeController.GetDetail)

		// http://127.0.0.1:8080/api/store/create
		storeAPI.Post("/create", storeController.Create)

		// http://127.0.0.1:8080/api/store/delete
		storeAPI.Post("/delete", storeController.Delete)

		// http://127.0.0.1:8080/api/store/update
		storeAPI.Post("/update", storeController.Update)
	}
}
