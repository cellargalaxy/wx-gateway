package controller

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/wx-gateway/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

//登录检查
func validate(c *gin.Context) {
	util.Validate(c, validateHandler)
}
func validateHandler(c *gin.Context) (string, jwt.Claims, error) {
	return config.Config.Secret, &jwt.StandardClaims{}, nil
}
