package service

import (
	"giftcard/application/repo"
)

type SystemDictType struct {
	repo *repo.SystemDictType
}

func NewSystemDictType() SystemDictType {
	service := SystemDictType{}
	service.repo = repo.NewSystemDictType()
	return service
}
