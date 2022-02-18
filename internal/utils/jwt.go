package utils

import (
	"github.com/dollarkillerx/fireworks/internal/conf"
	"github.com/dollarkillerx/jwt"
)

var JWT *jwt.JWT

func InitJwt(cfg conf.BackendConfig) {
	JWT = jwt.NewJwt(cfg.JWTToken)
}
