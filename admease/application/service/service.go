package service

import (
	"admease/application/dto"
	"github.com/micrease/gorme"
)

func ToSystemPage[T any](result *gorme.PageResult[T]) dto.SystemPageListResp[T] {
	var t dto.SystemPageListResp[T]
	t.PageInfo = dto.PageInfo{
		CurrentPage: result.PageNo,
		TotalPage:   result.TotalPage,
		Total:       result.TotalSize,
	}
	t.Items = result.List
	return t
}
