package controller

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg-gateway/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

//登录检查
func validate(ctx *gin.Context) {
	util.Validate(ctx, validateHandler)
}
func validateHandler(ctx *gin.Context) (string, jwt.Claims, error) {
	return config.Config.Secret, &jwt.StandardClaims{}, nil
}
