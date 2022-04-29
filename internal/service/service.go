package service

import (
	"github.com/Feride3d/banner-rotation-service/internal/storage"
)

type Service struct { //

}

func NewService(bd *storage.Storage) *Service { // сервис обращается к БД через указатель на структуру БД (dependency injunction)
	return &Service{}
}
