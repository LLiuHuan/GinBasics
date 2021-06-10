package global

import (
	"fairman-server/config"
	"github.com/docker/docker/client"
	ut "github.com/go-playground/universal-translator"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	FMS_VP     *viper.Viper
	FMS_CONFIG  config.Server
	FMS_LOG    *zap.Logger
	FMS_TRANS  ut.Translator
	FMS_CIL    *client.Client
)
