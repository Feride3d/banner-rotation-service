// Имплементируем хэндлеры
package handler

import (
	"github.com/Feride3d/banner-rotation-service/internal/service"
)

type Handler struct {
	services *service.Service // dependency injunction
}

func NewHandler(services *service.Service) *Handler { // обработчики вызывают методы сервиса
	return &Handler{services: services}
}
