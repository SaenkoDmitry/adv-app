package models

type AdvSearchReq struct {
	Page      string `form:"page" binding:"required"`
	Limit     string `form:"limit"`
	OrderBy   string `form:"orderBy" binding:"oneof=price created_at"`
	Direction string `form:"direction" binding:"oneof=asc desc"`
}
