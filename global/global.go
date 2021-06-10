package global

import (
	"gin-basics/config"
	"github.com/docker/docker/client"
	ut "github.com/go-playground/universal-translator"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	GB_VP     *viper.Viper
	GB_CONFIG config.Server
	GB_LOG    *zap.Logger
	GB_TRANS  ut.Translator
	GB_CIL    *client.Client
)
