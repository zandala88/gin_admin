package service

import "gin_admin/define"

func NewQueryRequest() *QueryRequest {
	return &QueryRequest{
		Page:    1,
		Size:    define.DefaultSize,
		Keyword: "",
	}
}
