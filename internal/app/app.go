package app

import (
	"github.com/UnitedIngvar/onmi_test/internal/config"
	"github.com/UnitedIngvar/onmi_test/internal/transport/rest/server"
)

func Run() {
	cfg := config.MustLoadConfig()
	log := config.SetupLogger(cfg.Env)

	server.StartServer(log, cfg)
}
