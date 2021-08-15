package app

import (
	"github.com/gin-gonic/gin"
	accesstoken "github.com/oauth/api/src/domain/access_token"
	newhttp "github.com/oauth/api/src/http"
	"github.com/oauth/api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {

	atHandler := newhttp.NewHandler(accesstoken.NewService(db.NewRepository()))
	router.GET("v1/oauth/access_token:access_token_id", atHandler.GetById)
	router.Run(":8091")
}
