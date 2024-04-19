package config

import "os"

type AdminConfig struct {
	AdminUsername string
	AdminPassword string
}

func loadAdminConfig() AdminConfig {
	return AdminConfig{
		AdminUsername: os.Getenv("ADMIN_USERNAME"),
		AdminPassword: os.Getenv("ADMIN_PASSWORD"),
	}
}
