package routers

import (
	"thinkdecideact/src/modules/store"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ApiRouter(routerGroup *gin.RouterGroup, db *gorm.DB) {
	store.Router(routerGroup, db)
}
