package config

import "os"

type AdminConfig struct {
	adminUsername string
	adminPassword string
}

func loadAdminConfig() AdminConfig {
	return AdminConfig{
		adminUsername: os.Getenv("ADMIN_USERNAME"),
		adminPassword: os.Getenv("ADMIN_PASSWORD"),
	}
}

func (adminConfig *AdminConfig) UserName() string {
	return adminConfig.adminUsername
}

func (adminConfig *AdminConfig) Password() string {
	return adminConfig.adminPassword
}
