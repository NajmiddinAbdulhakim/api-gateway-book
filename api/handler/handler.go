package handler

import (
	"github.com/NajmiddinAbdulhakim/ude/api-gateway/config"
	"github.com/NajmiddinAbdulhakim/ude/api-gateway/service"
)

type Handler struct {
	cfg            config.Config
	serviceManager service.IServiceManager
}

type HandlerConfig struct {
	Cfg            config.Config
	ServiceManager service.IServiceManager
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
	}
}
