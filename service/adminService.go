package service

import (
	"github.com/Puttipong1/assessment-tax/config"
	"github.com/Puttipong1/assessment-tax/server"
)

type AdminService struct {
	config *config.Config
}

func NewAdminService(server *server.Server) *AdminService {
	return &AdminService{
		config: server.Config,
	}
}
