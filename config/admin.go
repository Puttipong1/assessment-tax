package config

import (
	"os"

	"github.com/Puttipong1/assessment-tax/common"
)

type AdminConfig struct {
	adminUsername string
	adminPassword string
}

func loadAdminConfig() AdminConfig {
	//log := Logger()
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")
	if username == "" || password == "" {
		log.Fatal().Msgf(common.GetEnvErrorMessage, "ADMIN_USERNAME or ADMIN_PASSWORD", common.ShutDownServerMessage)
	}
	return AdminConfig{
		adminUsername: username,
		adminPassword: password,
	}
}

func (adminConfig *AdminConfig) UserName() string {
	return adminConfig.adminUsername
}

func (adminConfig *AdminConfig) Password() string {
	return adminConfig.adminPassword
}
