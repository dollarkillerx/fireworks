package utils

import (
	"github.com/dollarkillerx/fireworks/internal/conf"
	"github.com/dollarkillerx/jwt"
)

var JWT *jwt.JWT

func InitJWT() {
	JWT = jwt.NewJwt(conf.GetBackendConfig().JWTToken)
}
