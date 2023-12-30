package lambda

import (
	"learn-to-code/internal/infrastructure/config"
	"learn-to-code/internal/infrastructure/service"
)

type HandlerBase struct {
	Cfg               config.Config
	RegistryOverrides []service.RegistryOverride
}
