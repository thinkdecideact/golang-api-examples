package routers

import (
	"thinkdecideact/src/apps/store"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

func ApiRouter(db *gorm.DB) func(iris.Party) {
	return func(r iris.Party) {
		store.Router(r, db)
	}
}
