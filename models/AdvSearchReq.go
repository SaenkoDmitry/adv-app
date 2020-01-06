package models

type AdvSearchReq struct {
	Page      string `form:"page" binding:"required"`
	Limit     string `form:"limit" binding:"required"`
	OrderBy   string `form:"orderBy" binding:"oneof=name price"`
	Direction string `form:"direction" binding:"oneof=asc desc"`
}
