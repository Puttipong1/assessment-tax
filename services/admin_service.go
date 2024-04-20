package services

import (
	"net/http"

	"github.com/Puttipong1/assessment-tax/config"
	"github.com/Puttipong1/assessment-tax/models/request"
	"github.com/Puttipong1/assessment-tax/models/response"
	"github.com/Puttipong1/assessment-tax/server"
)

type AdminService struct {
	config     *config.Config
	JwtService *JwtService
}

func (adminService *AdminService) Login(login request.Login) (response.Token, int) {
	log := config.Logger()
	if login.Username != adminService.adminUsername() || login.Password != adminService.adminPassword() {
		log.Error().Msg("Username or Password incorrect")
		return response.Token{}, http.StatusUnauthorized
	}
	return adminService.JwtService.CreateToken()
}

func NewAdminService(server *server.Server) *AdminService {
	return &AdminService{
		config:     server.Config,
		JwtService: NewJwtService(server.Config)}
}

func (adminService *AdminService) adminUsername() string {
	return adminService.config.AdminConfig().UserName()
}

func (adminService *AdminService) adminPassword() string {
	return adminService.config.AdminConfig().Password()
}
